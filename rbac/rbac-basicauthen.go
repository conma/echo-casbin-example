package rbac

import (
	"github.com/casbin/casbin"
	"github.com/labstack/echo"
	"net/http"
)


func BasicAuthenServe() {
	e := echo.New()
	enforcer, _ := casbin.NewEnforcer("echo-casbin-example/rbac/rbac-basicauthen-model.conf", "echo-casbin-example/rbac/rbac-basicauthen-policy.csv")
	ef := BasicAuthenEnforcer{enforcer}
	e.Use(ef.BasicAuthenEnforcer)
	e.GET("/guest", GetFooPage)
	e.GET("/admin", GetBarPage)
	e.POST("/admin", PostBarPage)
	e.Logger.Fatal(e.Start("0.0.0.0:3000"))
}

func GetFooPage(c echo.Context) error {
	return c.JSON(http.StatusOK, "GET /guest OK")
}

func GetBarPage(c echo.Context) error {
	return c.JSON(http.StatusOK, "GET /admin OK")
}

func PostBarPage(c echo.Context) error {
	return c.JSON(http.StatusOK, "POST /admin OK")
}