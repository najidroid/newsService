package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int
	Title    string
	Link     string
	Desc     string
	ImageUri string
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(User))
}

func SetUsers() []*User {
	fmt.Print("**************hellllooooooo*******************")
	var data []*User
	orm.NewOrm().QueryTable(new(User)).All(&data)
	return data
}
