package routers

import (
	"github.com/najidroid/newsService/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//sedongsdg
	//test
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user", beego.NSInclude(&controllers.UserController{})),
	)
	beego.AddNamespace(ns)
}
