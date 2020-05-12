Go echo casbin rbac example: basic authen and loginform authen  
Required: echo, casbin  
install echo: `go get -u github.com/labstack/echo/`  
install casbin: `go get github.com/casbin/casbin`  
install this example: clone into $GOPATH/src
Run:
- basic authen: edit file $GOPATH/src/echo-casbin-example/main.go  
main.go:  
`//rbac.LoginServe()`  
`rbac.BasicAUthenServe()`  
 Test:  
 `curl guest:@localhost:3000/guest` => OK  
 `curl guest:@localhost:3000/admin` => Forbidden  
 `curl admin:@localhost:3000/admin` => OK
 - login with login form: edit file $GOPATH/src/echo-casbin-example/main.go  
 main.go:  
 `rbac.LoginServe()`  
 `//rbac.BasicAUthenServe()`  
 Test:  
 `curl localhost:3001/member` => Forbidden  
 `curl -X POST localhost:3001/login -H 'Content-Type: application/json' -d '{"name":"Admin"}'` => OK  
 `curl localhost:3001/admin` => OK  
 `curl -X POST localhost:3001/login -H 'Content-Type: application/json' -d '{"name":"Member1"}'` => OK  
 `curl localhost:3001/admin` => Forbidden  
 `curl localhost:3001/member` => OK