package main

import (
	"zhengbiwen/blog-server/cron"
	"zhengbiwen/blog-server/models"
	"zhengbiwen/blog-server/routers"
	"zhengbiwen/blog-server/session"
	"zhengbiwen/blog-server/utils"
)

func main() {
	db, err := models.InitDB()
	session.LoadSessionsFromDB()

	cron.Run()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := routers.InitRouter()
	r.Run(utils.HTTPPort)
}
