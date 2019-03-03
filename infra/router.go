package infra

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/tadoku/api/domain"
	"github.com/tadoku/api/interfaces/services"
	"github.com/tadoku/api/usecases"
)

// NewRouter instantiates a router
func NewRouter(
	port string,
	jwtSecret string,
	routes ...services.Route,
) services.Router {
	m := &middlewares{
		restrict:      newJWTMiddleware(jwtSecret),
		authenticator: usecases.NewRoleAuthenticator(),
	}
	e := newEcho(m, routes...)
	return router{e, port}
}

type middlewares struct {
	restrict      echo.MiddlewareFunc
	authenticator usecases.RoleAuthenticator
}

func newEcho(
	m *middlewares,
	routes ...services.Route,
) *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = errorHandler

	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			e.GET(route.Path, wrap(route, m))
		case http.MethodPost:
			e.POST(route.Path, wrap(route, m))
		default:
			log.Fatalf("HTTP verb %v is not supported", route.Method)
		}
	}

	return e
}

func newJWTMiddleware(secret string) echo.MiddlewareFunc {
	cfg := middleware.JWTConfig{
		Claims:     &jwtClaims{},
		SigningKey: []byte(secret),
	}
	return middleware.JWTWithConfig(cfg)
}

var errorCodeRegularExpression = regexp.MustCompile("^code=([0-9]{3}).")

func errorHandler(err error, c echo.Context) {
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

	c.NoContent(http.StatusInternalServerError)
}

func (m *middlewares) authenticateRole(c echo.Context, minRole domain.Role) error {
	u, err := (&context{c}).User()
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

		return r.HandlerFunc(&context{c})
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
