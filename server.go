package main

import (
	"fmt"
	"net/http"

	router "github.com/xiemenger/GolangMock/http"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
)

func main() {
	const port string = ":8080"
	httpRouter.GET("/", func(resp http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(resp, "Up and running.....")
	})
	httpRouter.GET("/posts", getPosts)
	httpRouter.POST("/posts", addPost)
	httpRouter.SERVE(port)

}
