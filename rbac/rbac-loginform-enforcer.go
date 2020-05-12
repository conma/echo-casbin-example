package rbac

import (
	"github.com/casbin/casbin"
	"github.com/labstack/echo"
)

type LoginFormEnforcer struct {
	enforcer *casbin.Enforcer
}

func (e *LoginFormEnforcer) LoginFormEnforcer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := User{}.GetLoginUserRole()
		path := c.Request().URL.Path
		act := c.Request().Method

		result, _ := e.enforcer.Enforce(role, path, act)

		if result {
			return next(c)
		}
		return echo.ErrForbidden
	}
}