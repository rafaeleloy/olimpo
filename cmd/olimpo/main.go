package main

import (
	"time"

	"olimpo/app/http/route"
	application "olimpo/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	app := application.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)
}
