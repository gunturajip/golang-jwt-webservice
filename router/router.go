package router

import (
	"golang-jwt-auth/controllers"
	"golang-jwt-auth/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Iki dikembangke ambek guntur wkwk")
	})

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.GET("/", controllers.GetProducts)
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.GET("/:productId", middlewares.ProductAuthorization(), controllers.GetProduct)
		productRouter.PUT("/:productId", middlewares.ProductAuthorization(), controllers.UpdateProduct)
		productRouter.DELETE("/:productId", middlewares.ProductAuthorization(), controllers.DeleteProduct)
	}

	return r
}
