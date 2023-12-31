package main

import (
	"log"

	"github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/app"
	helper "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/helpers"
)

func main() {
	/**
	@description Setup Server
	*/
	router := app.SetupRouter()
	/**
	@description Run Server
	*/
	log.Fatal(router.Run(":" + helper.GodotEnv("GO_PORT")))
}
