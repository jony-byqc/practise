package main

import (
	"github.com/kataras/iris/v12"
	//"golang.org/x/sync/errgroup"
)

//var g errgroup.Group
var sig string
var app2 *iris.Application

func main() {
	type boxDetail struct {
		fileName []string
		count    int
	}
	BoxDetail := boxDetail{
		//fileName:,
		count: 9,
	}
	app := iris.Default()
	app.Handle("GET", "/ping", func(context iris.Context) {
		context.JSON(iris.Map{"message": BoxDetail})
	})
	//监听端口
	app.Run(iris.Addr(":8089"))

}

//中间件
func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
