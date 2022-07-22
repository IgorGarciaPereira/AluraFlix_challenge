package controller

import(
  "net/http"

  "github.com/gin-gonic/gin"

  "aluraFlixAPI/database"
  "aluraFlixAPI/database/models"
)

func ListVideos(c *gin.Context){
  var videos []models.Video
  database.DB.Find(&videos)

  c.JSON(http.StatusOK, videos)
}

func GetVideo(c *gin.Context){
  id := c.Param("id")
  var video models.Video
  query := database.DB.First(&video, id)

  if query.Error != nil {
    c.JSON(http.StatusNotFound, nil)
    return
  }

  c.JSON(http.StatusOK, video)
}
