package routers

import (
	"web/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSCond(func(ctx *context.Context) bool {
			
		}),
	)
	
	beego.AddNamespace(ns)
}
