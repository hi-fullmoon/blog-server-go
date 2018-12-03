package main

import (
	"zhengbiwen/blog_management_system/models"
	"zhengbiwen/blog_management_system/routers"
)

func main() {
	db, err := models.InitDB()
	//session.LoadSessionsFromDB()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := routers.InitRouter()
	r.Run()
}
