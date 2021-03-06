package infra

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/tadoku/tadoku/services/identity-api/domain"
	"github.com/tadoku/tadoku/services/identity-api/interfaces/services"
	"github.com/tadoku/tadoku/services/identity-api/usecases"
)

// NewRouter instantiates a router
func NewRouter(
	environment domain.Environment,
	port string,
	jwtSecret string,
	sessionCookieName string,
	corsAllowedOrigins []string,
	errorReporter usecases.ErrorReporter,
	routes ...services.Route,
) services.Router {
	m := &middlewares{
		restrict:      newJWTMiddleware(jwtSecret, sessionCookieName),
		authenticator: usecases.NewRoleAuthenticator(),
	}
	e := newEcho(environment, m, corsAllowedOrigins, errorReporter, routes...)
	return router{e, port}
}

type middlewares struct {
	restrict      echo.MiddlewareFunc
	authenticator usecases.RoleAuthenticator
}

func newEcho(
	environment domain.Environment,
	m *middlewares,
	corsAllowedOrigins []string,
	errorReporter usecases.ErrorReporter,
	routes ...services.Route,
) *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = errorHandler(errorReporter)
	e.Use(newContextMiddleware(environment))
	e.Use(sentryecho.New(sentryecho.Options{}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: corsAllowedOrigins,
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	for _, route := range routes {
		e.Add(route.Method, route.Path, wrap(route, m))
	}

	return e
}

func newContextMiddleware(environment domain.Environment) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &context{c, environment}
			return next(cc)
		}
	}
}

func newJWTMiddleware(secret, sessionCookieName string) echo.MiddlewareFunc {
	cfg := middleware.JWTConfig{
		Claims:      &jwtClaims{},
		SigningKey:  []byte(secret),
		TokenLookup: fmt.Sprintf("cookie:%s", sessionCookieName),
	}
	return middleware.JWTWithConfig(cfg)
}

var errorCodeRegularExpression = regexp.MustCompile("^code=([0-9]{3}).")

func errorHandler(errorReporter usecases.ErrorReporter) func(error, echo.Context) {
	return func(err error, c echo.Context) {
		c.Logger().Error(err)

		if err == middleware.ErrJWTMissing {
			c.NoContent(http.StatusUnauthorized)
			return
		}

		if match := errorCodeRegularExpression.FindStringSubmatch(err.Error()); len(match) > 1 {
			if statusCode, errInt := strconv.Atoi(match[1]); errInt == nil {
				c.NoContent(statusCode)
				return
			}
		}

		if err == domain.ErrNotFound {
			c.NoContent(http.StatusNotFound)
			return
		}

		errorReporter.Capture(err)
		c.NoContent(http.StatusInternalServerError)
	}
}

func (m *middlewares) authenticateRole(c echo.Context, minRole domain.Role) error {
	u, err := c.(*context).User()
	if err == ErrEmptyUser && minRole != domain.RoleGuest {
		return c.NoContent(http.StatusUnauthorized)
	}
	err = m.authenticator.IsAllowed(u, minRole)

	if err != nil {
		return c.NoContent(http.StatusForbidden)
	}

	return nil
}

func wrap(r services.Route, m *middlewares) echo.HandlerFunc {
	handler := func(c echo.Context) error {
		err := m.authenticateRole(c, r.MinRole)
		if err != nil {
			return err
		}

		return r.HandlerFunc(c.(*context))
	}

	if r.MinRole > domain.RoleGuest {
		handler = m.restrict(handler)
	}

	return handler
}

type router struct {
	*echo.Echo
	port string
}

func (r router) StartListening() error {
	return r.Start(":" + r.port)
}
