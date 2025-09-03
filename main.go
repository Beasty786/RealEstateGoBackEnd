package main

import (

	"restate_backend/config"
	"restate_backend/router"

)

func main (){
	// Might wanna set up some environment stuff here
	port := "8080"
	// Set up gin server
	init := config.Init()

	app := router.Init(init)
	app.Run(":" + port) 

}