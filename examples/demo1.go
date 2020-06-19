package main

import (
	"fmt"

	"github.com/jiujuan/lilac"
)

func indexHandler(ctx *lilac.Context) {
	fmt.Println("hello world!")
	ctx.String(200, "HelloWorld!")
}

type Message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func helloHandler(ctx *lilac.Context) {
	name := ctx.Param("name")
	fmt.Println("hello handler: ", name)
	//ctx.String(200, name)
	ctx.JSON(200, Message{200, name})
}

func main() {
	app := lilac.New()

	app.Router.GET("/", indexHandler)
	app.Router.POST("/hello/:name", helloHandler)
	app.Run(":8080")
}
