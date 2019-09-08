package main

import (
	_ "clever-cloud.com/app_ef9c76b3-66e6-4a0f-8eb7-b645d18a05e1.git"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"

	"fmt"

	//	"github.com/astaxie/beego/orm"

	"log"

	"github.com/claudiu/gocron"

	//"encoding/xml"
	//	"io/ioutil"
	//	"net/http"

	"github.com/ungerik/go-rss"

	//	"github.com/go-telegram-bot-api/telegram-bot-api"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

//type User struct {
//	Id    int
//	Title string
//	Link  string
//	Desc  string
//}

func init() {
	//	orm.RegisterDriver("mysql", orm.DRMySQL)

	//	orm.RegisterDataBase("default", "mysql", "root:root@/newsservice?charset=utf8")
}

func main() {
	//	// Database alias.
	//	name := "default"

	//	// Drop table and re-create.
	//	force := true

	//	// Print log.
	//	verbose := true

	//	// Error.
	//	err := orm.RunSyncdb(name, force, verbose)

	//	if err != nil {
	//		fmt.Println(err)

	//	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//	orm.RegisterModel(new(User))

	readRSS()

	//	startBot()

	//startAnotherBot()

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
	gocron.Every(1).Minutes().Do(readRSS)
	//gocron.Every(10).Seconds().Do(vijay)
	//gocron.Every(1).Monday().Do(task)
	//gocron.Every(1).Thursday().At("18:30").Do(doTownCup)
	//gocron.Every(1).Friday().At("18:30").Do(doAlleyCup)

	s.Start()
}

func readRSS() {
	fmt.Println("reading RSS *******************************************")
	//	channel, err := rss.Read("https://www.irinn.ir/fa/rss/allnews")
	channel, err := rss.Read("https://www.irinn.ir/fa/rss/1")
	//	channel, err := rss.Read("https://www.varzesh3.com/rss/all")
	//	channel, err := rss.Read("https://www.tasnimnews.com/fa/rss/feed/0/8/0/%D9%85%D9%87%D9%85%D8%AA%D8%B1%DB%8C%D9%86-%D8%A7%D8%AE%D8%A8%D8%A7%D8%B1-%D8%AA%D8%B3%D9%86%DB%8C%D9%85")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(channel.Title)

	b, err := tb.NewBot(tb.Settings{
		Token: "592949403:AAG-CkEkdqZYxN6DcPGVv8dzAErzIwxNLWQ",
		// You can also set custom API URL. If field is empty it equals to "https://api.telegram.org"
		//		URL: "http://195.129.111.17:8012",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	rec := &tb.Chat{ID: -1001212999492, Type: "channel", FirstName: "test", Username: "thisistestchann"}

	for _, item := range channel.Item {
		text := item.Title + "\n" + item.Description + "\n" + "لینک خبر: " + item.Link

		fmt.Println(item.Title)
		fmt.Println(item.Link)
		fmt.Println(item.Description)
		if item.Enclosure != nil {
			fmt.Println(item.Enclosure[0].URL)
			pic := &tb.Photo{File: tb.FromURL(item.Enclosure[0].URL), Caption: text}
			b.Send(rec, pic)
		} else {
			b.Send(rec, text)
		}

	}

	b.Start()
}

//func startAnotherBot() {
//	b, err := tb.NewBot(tb.Settings{
//		Token: "592949403:AAG-CkEkdqZYxN6DcPGVv8dzAErzIwxNLWQ",
//		// You can also set custom API URL. If field is empty it equals to "https://api.telegram.org"
//		//		URL: "http://195.129.111.17:8012",
//		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
//	})

//	if err != nil {
//		log.Fatal(err)
//		return
//	}

//	b.Handle("/hello", func(m *tb.Message) {
//		fmt.Println("11111111111111111111111")
//		b.Send(m.Sender, "hello world")
//	})

//	b.Handle(tb.OnChannelPost, func(m *tb.Message) {
//		if m == nil {
//			fmt.Println("m is nil*********************")
//		}
//		//		fmt.Println(tb.Recipient())
//		fmt.Println(m.Chat)
//		//		p := &tb.Photo{File: tb.FromDisk("nazar.jpg")}
//		//		b.Send(m.Chat, p)
//		//b.Send(m.Sender, "hello world")
//		// channel posts only
//	})

//	rec := &tb.Chat{ID: -1001212999492, Type: "channel", FirstName: "test", Username: "thisistestchann"}
//	p1 := &tb.Photo{File: tb.FromDisk("nazar.jpg")}
//	p2 := &tb.Photo{File: tb.FromURL("https://www.google.com/imgres?imgurl=https%3A%2F%2Fimages.unsplash.com%2Fphoto-1472214103451-9374bd1c798e%3Fixlib%3Drb-1.2.1%26ixid%3DeyJhcHBfaWQiOjEyMDd9%26w%3D1000%26q%3D80&imgrefurl=https%3A%2F%2Funsplash.com%2Fsearch%2Fphotos%2Fpic&docid=yLbmyYyJ2Ux6uM&tbnid=igUvhVcxMorBOM%3A&vet=10ahUKEwjKlcX2yrnkAhVok4sKHTQDDocQMwhtKAAwAA..i&w=1000&h=667&client=ubuntu&bih=639&biw=1299&q=pic&ved=0ahUKEwjKlcX2yrnkAhVok4sKHTQDDocQMwhtKAAwAA&iact=mrc&uact=8")}
//	b.SendAlbum(rec, tb.Album{p1, p2})

//	p := &tb.Photo{File: tb.FromDisk("nazar.jpg"), Caption: "این یک کپشن است:)"}
//	b.Send(rec, p)

//	b.Send(rec, "محمد درستش کردم، همه چیز رو یاد گرفتم @mohammad_cs_15")

//	b.Start()

//}
