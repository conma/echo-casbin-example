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
		var result bool
		role := User{}.GetLoginUserRole()
		path := c.Request().URL.Path
		act := c.Request().Method

		// handle for delete posts
		if path == "/post/delete" {
			if role == "member" {
				loginName := User{}.GetLoginName()
				post := Post{}
				c.Bind(post)
				if loginName != post.Author {
					return echo.ErrForbidden
				}
			}
		}

		result, _ = e.enforcer.Enforce(role, path, act)

		if result {
			return next(c)
		}
		return echo.ErrForbidden
	}
}
