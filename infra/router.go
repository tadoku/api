package infra

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/tadoku/api/domain"
	"github.com/tadoku/api/interfaces/services"
)

// NewRouter instantiates a router
func NewRouter(
	port string,
	jwtSecret string,
	routes ...services.Route,
) services.Router {
	e := newEcho(jwtSecret, routes...)
	return router{e, port}
}

func newEcho(jwtSecret string, routes ...services.Route) *echo.Echo {
	e := echo.New()
	restricted := newJWTMiddleware(jwtSecret)

	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			e.GET(route.Path, wrap(route, restricted))
		case http.MethodPost:
			e.POST(route.Path, wrap(route, restricted))
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

func isRoleAllowed(c echo.Context, minRole domain.Role) bool {
	ctx := &context{c}
	u, err := ctx.User()
	if err != nil {
		return false
	}

	if u.Role < minRole {
		return false
	}

	return true
}

func wrap(r services.Route, restrict echo.MiddlewareFunc) echo.HandlerFunc {
	handler := func(c echo.Context) error {
		// @TODO: find out if we can do this nicer with a middleware
		if !isRoleAllowed(c, r.RoleRestriction) {
			return c.NoContent(http.StatusForbidden)
		}

		return r.HandlerFunc(&context{c})
	}

	if r.Restricted {
		handler = restrict(handler)
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
