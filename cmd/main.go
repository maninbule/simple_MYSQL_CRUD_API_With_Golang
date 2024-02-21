package main

import (
	"fmt"
	"github.com/simple_MYSQL_CRUD_API_With_Golang/pkg/routes"
	"net/http"
)

func main() {
	routerBook := routes.CreateRouter()
	fmt.Println("start server on port:8080")
	err := http.ListenAndServe(":8080", routerBook)
	panic(err)
}
