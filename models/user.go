package models

import (
	"fmt"

	//	"github.com/astaxie/beego/orm"
)

type User struct {
	Id    int
	Title string
	Link  string
	Desc  string
}

func init() {
	// Need to register model in init
	//	orm.RegisterModel(new(User))
}

func SetUsers() {
	fmt.Print("**************hellllooooooo*******************")
}
