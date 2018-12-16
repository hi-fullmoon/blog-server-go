package main

import (
	"zhengbiwen/blog-server/models"
	"zhengbiwen/blog-server/routers"
	"zhengbiwen/blog-server/session"
)

func main() {
	db, err := models.InitDB()
	session.LoadSessionsFromDB()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := routers.InitRouter()
	r.Run()
}
