package main

import (
	"golang-todo-api/config"
	"golang-todo-api/handlers"
	"golang-todo-api/repositories"
	"golang-todo-api/routes"
	"golang-todo-api/services"
)

func main() {
	env := config.LoadEnv()
	db := config.ConnectDB(env)
	config.AutoMigrate(db)
	app := config.NewFiberApp(env)
	respository := repositories.NewTodoRepository(db)
	service := services.NewTodoService(respository)
	handler := handlers.NewTodoHandler(service)
	routes.LoadRoutes(app, handler)
	config.ListenFiberApp(app, env)
}
