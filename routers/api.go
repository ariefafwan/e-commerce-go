package routers

import (
	"e-commerce-go/internal/controllers"
	"e-commerce-go/internal/repositories"
	"e-commerce-go/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouters(router *gin.Engine, db *gorm.DB) {

	userRepo := repositories.NewuserRepository(db)
	authController := controllers.NewAuthController(userRepo)

	router.POST("/login", authController.Login)
	router.POST("/register", authController.Register)
	router.POST("/refresh", authController.Refresh)
	router.POST("/logout", authController.Logout)

	api := router.Group("/api")
	api.Use(middleware.JWTMiddleware())
	{
		tokoRepo := repositories.NewMasterTokoRepository(db)
		tokoController := controllers.NewMasterTokoController(tokoRepo)
		toko := api.Group("/master-toko")
		{
			toko.GET("/", tokoController.GetToko)
			toko.PUT("/:id", tokoController.UpdateToko)
		}
	}
}
