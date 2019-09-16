package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type UserIsna struct {
	Id         int
	Title      string
	Link       string
	Desc       string `orm:"size(1000)"`
	ImageUri   string
	Type       string
	PubDate    time.Time
	PubDateStr string
}

type UserIsnaKhabardar struct {
	Id         int
	Title      string
	Link       string
	Desc       string `orm:"size(1000)"`
	ImageUri   string
	Type       string
	PubDate    time.Time
	PubDateStr string
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
