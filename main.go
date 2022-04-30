package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"project_1/controller"
	"project_1/utils"
)

var (
	config *utils.Config
)

func init() {
	var err error
	config, err = utils.ParseConfig("./config/app.json")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	app := gin.Default()

	controller.Router(app)

	app.Run(config.AppHost + ":" + config.AppPort)
}
