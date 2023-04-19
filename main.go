package main

import (
	"gin-example/repo"
	"gin-example/router"
	"gin-example/util"
)

func main() {
	repo.InitDB()
	defer repo.CloseDB()
	r := router.SetupRouter()
	r.Run(":8080")
	util.Logger().Info("Starting server...")
}
