package router

import (
	"github.com/gin-gonic/gin"

	"aluraFlixAPI/controller"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/", controller.Hello)

	video := r.Group("/videos")
	video.GET("/", controller.ListVideos)
	video.GET("/:id", controller.GetVideo)
	video.POST("/", controller.CreateVideo)
	video.PATCH("/:id", controller.UpdateVideo)
	video.DELETE("/:id", controller.DeleteVideo)

	r.Run()
}
