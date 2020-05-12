package rbac

import (
	"github.com/casbin/casbin"
	"github.com/labstack/echo"
)

type BasicAuthenEnforcer struct {
	enforcer *casbin.Enforcer
}

func (e *BasicAuthenEnforcer) BasicAuthenEnforcer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, _, _ := c.Request().BasicAuth()
		method := c.Request().Method
		path := c.Request().URL.Path

		result, _ := e.enforcer.Enforce(user, path, method)

		if result {
			return next(c)
		}
		return echo.ErrForbidden
	}
}