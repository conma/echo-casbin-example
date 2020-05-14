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
	Post struct {
		Id		string	`json:"id"`
		Content string	`json:"content"`
		Author	string	`json:"author"`
	}
)

var (
	users	map[string]User
	CurrentLoginUser User
	Posts map[string]Post
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
	e.GET("/post/list", GetPostsPage)
	e.POST("/post/delete", DeletePost)
	e.Logger.Fatal(e.Start("0.0.0.0:3001"))
}

func InitUsers() {
	users = map[string]User{}
	users["Admin"] = User{Name: "Admin", Role: "admin"}
	users["Member1"] = User{Name: "Member1", Role: "member"}
	users["Member2"] = User{Name: "Member2", Role: "member"}
}

func InitPosts() {
	Posts = map[string]Post{}
	Posts["1"] = Post{Id:"post1", Content:"Post1-Content", Author:"Member1"}
	Posts["2"] = Post{Id:"post2", Content:"Post2-Content", Author:"Member2"}
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

func (u User) GetLoginName() string {
	return CurrentLoginUser.Name
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
	InitPosts()
	return c.JSON(http.StatusOK, loginUser)
}

func MemberPage(c echo.Context) error {
	return c.JSON(http.StatusOK, "GET /member OK")
}

func AdminPage(c echo.Context) error {
	return c.JSON(http.StatusOK, "GET /admin OK")
}
func GetPostsPage(c echo.Context) error {
	return c.JSON(http.StatusOK, Posts)
}

func DeletePost(c echo.Context) error {
	p := new(Post)
	if  err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, "Post do not exists")
	}
	post := Posts[p.Id]
	if post.Id == "" {
		return c.JSON(http.StatusBadRequest, "Post do not exists")
	}
	delete(Posts, post.Id)
	return c.JSON(http.StatusOK, post)
}