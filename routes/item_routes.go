package routes

import (
	"aplikasi-daftar-belanja/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterItemRoutes(r *gin.Engine) {
	protected := r.Group("/")
	protected.Use(controllers.AuthMiddleware())
	{
		protected.GET("/items", controllers.GetItems)
		protected.POST("/items", controllers.CreateItem)
		protected.PATCH("/items/:id", controllers.UpdateItem)
		protected.DELETE("/items/:id", controllers.DeleteItem)
	}

}