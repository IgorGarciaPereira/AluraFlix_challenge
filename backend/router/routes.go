package router

import (
	"github.com/gin-gonic/gin"

	"aluraFlixAPI/controller"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/", controller.Hello)
	r.Run()
}
