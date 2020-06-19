package main

import (
	"github.com/jiujuan/lilac"
)

func UserIndex(ctx *lilac.Context) {
	ctx.String(200, "group router")
}

func GetUsername(ctx *lilac.Context) {
	ctx.String(200, "username : tom")
}

func Index(ctx *lilac.Context) {
	ctx.String(200, "main index")
}

func MiddlewareUser(ctx *lilac.Context) {
	ctx.String(200, "Middleware User, Yes!")
}

func main() {
	app := lilac.New()
	app.Router.GET("/index", Index)

	userGroupRouter := app.Group.Group("/user")
	{
		userGroupRouter.Use(MiddlewareUser)
		userGroupRouter.GET("/index", UserIndex)
		userGroupRouter.GET("/username", GetUsername)
	}

	app.Run(":8080")
}
