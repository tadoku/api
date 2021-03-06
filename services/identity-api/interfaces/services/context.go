//go:generate gex mockgen -source=context.go -package services -destination=context_mock.go

package services

import (
	"net/http"

	"github.com/tadoku/tadoku/services/identity-api/domain"
	"github.com/tadoku/tadoku/services/identity-api/usecases"
)

// based on https://github.com/labstack/echo/blob/a2d4cb9c7a629e2ee21861501690741d2374de10/context.go

// Context is a subset of the echo framework context, so we are not directly depending on it
type Context interface {
	// QueryParam returns the query param for the provided name.
	QueryParam(name string) string

	// IntQueryParam returns the query param for the provided name, converted to int
	IntQueryParam(name string) (uint64, error)

	// OptionalIntQueryParam returns the query param for the provided name, converted to int with a fallback
	OptionalIntQueryParam(name string, defaultValue uint64) (result uint64)

	// Get retrieves data from the context.
	Get(key string) interface{}

	// Set saves data in the context.
	Set(key string, val interface{})

	// Bind binds the request body into provided type `i`. The default binder
	// does it based on Content-Type header.
	Bind(i interface{}) error

	// String sends a string response with status code.
	String(code int, s string) error

	// NoContent sends a response with no body and a status code.
	NoContent(code int) error

	// JSON sends a JSON response with status code.
	JSON(code int, i interface{}) error

	// Claims gets all the user Claims
	Claims() *usecases.SessionClaims

	// User gets the current logged in user
	User() (*domain.User, error)

	// GetID gets the current id in the route
	GetID() (uint64, error)

	// Parses out the id in the route and binds it to the given variable
	BindID(*uint64) error

	// Sets a new cookie to send back to the client
	SetCookie(*http.Cookie)

	// Returns the environment the app is running in
	Environment() domain.Environment
}
