package controllers

import (
	//"encoding/json"
	//"fmt"
	"newsService/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// @Title Get
// @Description get all Users
// @Success 200 {object} models.User
// @router / user [get]
func (u *UserController) GetAll() {
	models.SetUsers()
	//models.SetUsers()
	//	u.Data["json"] = users
	//	u.ServeJSON()
}
