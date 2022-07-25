package controller

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "gopkg.in/validator.v2"

  "aluraFlixAPI/database"
  "aluraFlixAPI/database/models"
)

func CreateCategory(c *gin.Context){
  var category models.Category

  err := c.BindJSON(&category)
  if err != nil {
    c.JSON(http.StatusBadRequest, err)
    return
  }

  if err := validator.Validate(category); err != nil{
    c.JSON(http.StatusUnprocessableEntity, err)
    return
  }

  err = database.DB.Create(&category).Error;
  if err != nil {
    c.JSON(http.StatusBadRequest, err)
    return
  }
  c.JSON(http.StatusCreated, gin.H{
    "id": category.ID,
  })
}

func ListCategory(c *gin.Context){
  var categories []models.Category
  database.DB.Find(&categories);
  c.JSON(http.StatusOK, categories)
}

func GetCategory(c *gin.Context){
  var category models.Category
  id := c.Param("id")

  database.DB.First(&category, id)

  if category.ID == 0 {
    c.JSON(http.StatusNotFound, nil)
    return
  }

  c.JSON(http.StatusOK, category)
}

func UpdateCategory(c *gin.Context){
  var categoryInfo models.Category
  var category models.Category

  id := c.Param("id")

  if err := c.BindJSON(&categoryInfo); err != nil {
    c.JSON(http.StatusBadRequest, err)
    return
  }

  database.DB.First(&category, id)

  if category.ID == 0 {
    c.JSON(http.StatusNotFound, nil)
    return;
  }

  category.Title = categoryInfo.Title
  category.Color = categoryInfo.Color

  query := database.DB.Save(&category)
  if query.Error != nil {
    c.JSON(http.StatusInternalServerError, nil)
    return
  }

  c.JSON(http.StatusOK, category)
}

func DeleteCategory(c *gin.Context){
  var category models.Category
  id := c.Param("id")

  database.DB.First(&category, id)
  if category.ID == 0 {
    c.JSON(http.StatusNotFound, gin.H{
      "message": "Category not found",
    })
    return
  }

  err := database.DB.Delete(&category).Error
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "message": err,
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "message": "Category deleted successfully",
  })
}
