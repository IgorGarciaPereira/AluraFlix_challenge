package controller

import(
  "net/http"
  "strconv"

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

func CreateVideo(c *gin.Context){
  var video models.Video

  err := c.BindJSON(&video)
  if err != nil {
    c.JSON(http.StatusBadRequest, nil)
    return
  }

  database.DB.Create(&video)
  c.JSON(http.StatusCreated, gin.H{
    "id": video.ID,
  })
}

func UpdateVideo(c *gin.Context){
  var video models.Video
  var updateInfo models.Video

  idParam := c.Param("id")
  id, err := strconv.Atoi(idParam)

  if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{
        "message": "ID with incorrect type",
      })
      return
  }

  err = c.BindJSON(&updateInfo);
  if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{
        "message": "Data with incorrect type",
      })
      return
  }

  database.DB.First(&video, id)
  notExists := (video.ID == 0)
  if notExists {
    c.JSON(http.StatusNotFound, nil)
    return
  }

  video.Title = updateInfo.Title
  video.Description = updateInfo.Description
  video.Url = updateInfo.Url

  query := database.DB.Save(&video);
  if query.Error != nil {
    c.JSON(http.StatusInternalServerError, nil)
    return
  }

  c.JSON(http.StatusNoContent, nil)
}

func DeleteVideo(c *gin.Context){

  var video models.Video
  idParam := c.Param("id")
  id, err := strconv.Atoi(idParam)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "message": "ID incorrect format",
    })
    return
  }

  err = database.DB.Delete(&video, id).Error

  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "message": "error to delete video",
    })
  }

  c.JSON(http.StatusNoContent, nil)
}
