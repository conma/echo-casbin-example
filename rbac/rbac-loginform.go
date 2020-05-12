package rbac

import (
	"github.com/casbin/casbin"
	"github.com/labstack/echo"
	"net/http"
)

type (
	User struct {
		Name	string	`json:"name"`
		Role	string	`json:"role"`
	}
)

var (
	users	map[string]User
	CurrentLoginUser User
)

func LoginServe() {
	InitUsers()
	e := echo.New()
	enforcer, _ := casbin.NewEnforcer("echo-casbin-example/rbac/rbac-loginform-model.conf", "echo-casbin-example/rbac/rbac-loginform-policy.csv")
	ef := LoginFormEnforcer{enforcer}
	e.Use(ef.LoginFormEnforcer)
	e.POST("/login", Login)
	e.GET("/member", MemberPage)
	e.GET("/admin", AdminPage)
	e.Logger.Fatal(e.Start("0.0.0.0:3001"))
}

func InitUsers() {
	users = map[string]User{}
	users["Admin"] = User{Name: "Admin", Role: "admin"}
	users["Member1"] = User{Name: "Member1", Role: "member"}
	users["Member2"] = User{Name: "Member2", Role: "member"}
}

func (u User) GetUserByName(name string) User {
	return users[name]
}

func (u User) GetLoginUserRole() string {
	if CurrentLoginUser.Name == "" {
		return "guest"
	}
	return CurrentLoginUser.Role
}

func Login(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "User do not exists")
	}
	loginUser := User{}.GetUserByName(user.Name)
	if loginUser.Name == "" {
		return c.JSON(http.StatusBadRequest, "User do not exists")
	}
	CurrentLoginUser = User{Name: loginUser.Name, Role: loginUser.Role}
	return c.JSON(http.StatusOK, loginUser)
}

func MemberPage(c echo.Context) error {
	return c.JSON(http.StatusOK, "GET /member OK")
}

func AdminPage(c echo.Context) error {
	return c.JSON(http.StatusOK, "GET /admin OK")
}