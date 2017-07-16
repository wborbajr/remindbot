package main

import (
	"fmt"

	"github.com/aranair/remindbot/commands"
	"github.com/aranair/remindbot/config"
	"github.com/aranair/remindbot/handlers"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/BurntSushi/toml"
	"github.com/jasonlvhit/gocron"
)

func checkDues(ac handlers.AppContext) {
	// ac.SendText(chatId, text)
}

func main() {
	var conf config.Config

	_, err := toml.DecodeFile("configs.toml", &conf)
	checkErr(err)

	fmt.Println(conf)
	db := initDB(conf.DB.Datapath)
	defer db.Close()

	ac := handlers.NewAppContext(db, conf, commands.NewCommandList())
	// chatId := -6894201
	gocron.Every(5).Minutes().Do(checkDues, ac)
	fmt.Println("Starting timer")
	<-gocron.Start()
}

func initDB(datapath string) *sql.DB {
	db, err := sql.Open("sqlite3", datapath+"/reminders.db")
	checkErr(err)
	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
