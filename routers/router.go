package routers

import (
	"github.com/najidroid/newsService/controllers"

	"github.com/astaxie/beego"

	//	"fmt"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user", beego.NSInclude(&controllers.UserController{})),
	)
	beego.AddNamespace(ns)
}
