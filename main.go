package main

import "echo-casbin-example/rbac"

func main() {
	rbac.LoginServe()
	rbac.BasicAuthenServe()
}
