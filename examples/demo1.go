package main

import (
	"fmt"

	"github.com/jiujuan/lilac"
)

func indexHandler(ctx *lilac.Context) {
	fmt.Println("hello world!")
	ctx.String(200, "HelloWorld!")
}

func main() {
	app := lilac.New()

	app.GET("/", indexHandler)
	app.Run(":8080")
}
