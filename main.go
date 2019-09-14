package main

import (
	_ "github.com/najidroid/newsService/routers"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"

	"fmt"

	"github.com/astaxie/beego/orm"

	"log"

	"github.com/claudiu/gocron"

	"github.com/ungerik/go-rss"

	tb "gopkg.in/tucnak/telebot.v2"

	"github.com/najidroid/newsService/models"
)

var (
	bot *tb.Bot
	rec *tb.Chat
)

func init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "ua4bq61zbvkmnsrg:d7tZPzypUxp88hPKdcPk@tcp(bet6wf9aiup7rp3qths5-mysql.services.clever-cloud.com:3306)/bet6wf9aiup7rp3qths5?charset=utf8")

}

func main() {
	// Database alias.
	name := "default"

	// Drop table and re-create.
	force := true

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(name, force, verbose)

	if err != nil {
		fmt.Println(err)

	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	bot, _ = tb.NewBot(tb.Settings{
		Token: "592949403:AAG-CkEkdqZYxN6DcPGVv8dzAErzIwxNLWQ",
		// You can also set custom API URL. If field is empty it equals to "https://api.telegram.org"
		//		URL: "http://195.129.111.17:8012",
		//		Poller: &tb.LongPoller{Timeout: 1000 * time.Second},
	})

	rec = &tb.Chat{ID: -1001212999492, Type: "channel", FirstName: "test", Username: "thisistestchann"}

	readRSS()

	startGocorn()

	beego.Run()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		fmt.Println(err)
		fmt.Println(msg)
	}
}

func startGocorn() {
	gocron.Start()
	s := gocron.NewScheduler()
	gocron.Every(30).Minutes().Do(readRSS)
	//	gocron.Every(5).Seconds().Do(readRSS)
	//gocron.Every(1).Monday().Do(task)
	//gocron.Every(1).Thursday().At("18:30").Do(doTownCup)
	//gocron.Every(1).Friday().At("18:30").Do(doAlleyCup)

	s.Start()
}

func readRSS() {
	Isna("https://www.isna.ir/rss")

}

func Isna(url string) {
	channel, err := rss.Read(url)
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range channel.Item {
		text := item.Title + "\n" + item.Description + "\n" + "لینک خبر: " + item.Link

		if item.Enclosure != nil {
			pic := &tb.Photo{File: tb.FromURL(item.Enclosure[0].URL), Caption: text}
			bot.Send(rec, pic)

		} else {
			bot.Send(rec, text)
			//			this := models.UserIsna{Title: item.Title, Link: item.Link, Desc: item.Description, ImageUri: ""}
			//			_, err := orm.NewOrm().Insert(&this)
			//			if err != nil {
			//				fmt.Printf("save err... %s", err)
			//			}
		}
		this := models.UserIsna{Title: item.Title, Link: item.Link, Desc: item.Description, ImageUri: item.Enclosure[0].URL}
		_, err := orm.NewOrm().Insert(&this)
		if err != nil {
			fmt.Printf("save err... %s", err)
		}
	}
}
