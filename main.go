package main

import (
	"golang-todo-api/config"
)

func main() {
	env := config.LoadEnv()
	db := config.ConnectDB(env)
	config.AutoMigrate(db)
	config.NewFiberApp(env)
}
