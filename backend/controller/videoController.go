package controller

import(
  "net/http"
  "strconv"

  "github.com/gin-gonic/gin"
  "gopkg.in/validator.v2"

  "aluraFlixAPI/database"
  "aluraFlixAPI/database/models"
)

func ListVideos(c *gin.Context){
  var videos []models.Video

  search := c.Query("search")

  if search != "" {
    searchParams := "%" + search + "%"
    database.DB.Where("title LIKE ?", searchParams).Find(&videos)
  }else{
    database.DB.Find(&videos)
  }

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
  var category models.Category

  err := c.BindJSON(&video)
  if err != nil {
    c.JSON(http.StatusBadRequest, nil)
    return
  }



  if video.CategoryID == 0{
    video.CategoryID = 1
  }else{
    database.DB.First(&category, video.CategoryID)

    if category.ID == 0 {
      c.JSON(http.StatusNotFound, gin.H{
        "message": "category not found",
      })
      return
    }
  }

  if errs :=validator.Validate(video); errs != nil {
    c.JSON(http.StatusUnprocessableEntity, gin.H{
      "message": errs,
    })
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
