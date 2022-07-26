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

	category := r.Group("/category")
	category.GET("/", controller.ListCategory)
	category.GET("/:id/videos", controller.GetFullCategory)
	category.GET("/:id", controller.GetCategory)
	category.POST("/", controller.CreateCategory)
	category.PATCH("/:id", controller.UpdateCategory)
	category.DELETE("/:id", controller.DeleteCategory)

	r.Run()
}
