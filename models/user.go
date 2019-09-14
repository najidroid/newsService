package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type UserIsna struct {
	Id       int
	Title    string
	Link     string
	Desc     string
	ImageUri string
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(UserIsna))
}

func SetUsers() []*UserIsna {
	fmt.Print("**************hellllooooooo*******************")
	var data []*UserIsna
	orm.NewOrm().QueryTable(new(UserIsna)).All(&data)
	return data
}
