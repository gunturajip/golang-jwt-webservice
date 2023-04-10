package controllers

import (
	"golang-jwt-auth/database"
	"golang-jwt-auth/helpers"
	"golang-jwt-auth/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	db := database.GetDB()

	Product := []models.Product{}

	err := db.Preload("User").Find(&Product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully get products",
		"data":    Product,
	})
}

func GetProduct(c *gin.Context) {
	db := database.GetDB()

	Product := models.Product{}
	productId := c.Param("productId")

	err := db.Preload("User").First(&Product, "id = ?", productId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully get product",
		"data":    Product,
	})
}

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	Product := models.Product{}
	userId := uint(userData["id"].(float64))
	User := models.User{}
	errA := db.First(&User, "id = ?", userId).Error
	if errA != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": errA.Error(),
		})
		return
	}

	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userId
	Product.User = &User

	err := db.Create(&Product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully created product",
		"data":    Product,
	})
}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()

	Product := models.Product{}
	productId := c.Param("productId")

	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	err := db.Model(&Product).Where("id = ?", productId).Updates(models.Product{Title: Product.Title, Description: Product.Description}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully updated product",
	})
}

func DeleteProduct(c *gin.Context) {
	db := database.GetDB()

	Product := models.Product{}
	productId := c.Param("productId")

	err := db.Where("id = ?", productId).Delete(&Product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted product",
	})
}
