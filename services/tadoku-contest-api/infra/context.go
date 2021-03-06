package infra

import (
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/srvc/fail"

	"github.com/tadoku/tadoku/services/tadoku-contest-api/domain"
	"github.com/tadoku/tadoku/services/tadoku-contest-api/usecases"
)

// ErrEmptyUser when context contains no user when trying to get one
var ErrEmptyUser = fail.Errorf("user is empty")

type context struct {
	echo.Context
	environment domain.Environment
}

func (c context) Environment() domain.Environment {
	return c.environment
}

func (c context) User() (*domain.User, error) {
	claims := c.Claims()
	if claims != nil && claims.User != nil {
		return claims.User, nil
	}

	return nil, ErrEmptyUser
}

func (c context) Claims() *usecases.SessionClaims {
	if token, ok := c.Get("user").(*jwt.Token); ok {
		claims := token.Claims.(*jwtClaims)
		if claims != nil {
			return &claims.SessionClaims
		}
	}
	return nil
}

func (c context) GetID() (uint64, error) {
	idFromRoute := c.Param("id")
	id, err := strconv.ParseUint(idFromRoute, 10, 64)

	return id, domain.WrapError(err)
}

func (c context) BindID(id *uint64) error {
	value, err := c.GetID()
	*id = value

	return domain.WrapError(err)
}

func (c context) IntQueryParam(name string) (uint64, error) {
	param := c.QueryParam(name)
	result, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return 0, domain.WrapError(err)
	}

	return result, nil
}

func (c context) OptionalIntQueryParam(name string, defaultValue uint64) uint64 {
	result, err := c.IntQueryParam(name)
	if err != nil {
		return defaultValue
	}

	return result
}
