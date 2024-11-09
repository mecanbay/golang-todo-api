package config

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func NewFiberApp(env *Env) *fiber.App {
	preforkValue, err := strconv.ParseBool(env.FIBER_PREFORK)
	if err != nil {
		log.Println("FIBER_PREFORK only accepts a bool data type.")
	}

	config := fiber.Config{
		AppName:      env.FIBER_APPNAME,
		ServerHeader: env.FIBER_SERVER_HEADER,
		Prefork:      preforkValue,
	}

	app := fiber.New(config)

	return app
}

func ListenFiberApp(app *fiber.App, env *Env) {
	app.Listen(fmt.Sprintf("%v:%v", env.FIBER_HOST, env.FIBER_PORT))
}
