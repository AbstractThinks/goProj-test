package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"goProj-test/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})

	beego.Get("/restful/get", func(ctx *context.Context) {
		ctx.Output.Body([]byte("restful api get"))
	})

	beego.Post("/restful/post", func(ctx *context.Context) {
		ctx.Output.Body([]byte("bob"))
	})
}
