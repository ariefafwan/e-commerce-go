package routers

import (
	"e-commerce-go/internal/controllers"
	"e-commerce-go/internal/repositories"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func SetupRouters(router *gin.Engine, db *gorm.DB) {

		tokoRepo := repositories.NewMasterTokoRepository(db)
		tokoController := controllers.NewMasterTokoController(tokoRepo)

		toko := router.Group("/api")
		{
			toko.GET("/", tokoController.GetToko)
			toko.PUT("/:id", tokoController.UpdateToko)
		}
}
