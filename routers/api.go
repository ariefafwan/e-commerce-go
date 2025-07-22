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
	{
		tokoRepo := repositories.NewMasterTokoRepository(db)
		tokoController := controllers.NewMasterTokoController(tokoRepo)
		
		toko := api.Group("/master-toko")
		toko.Use(middleware.AuthMiddleware("Admin"))
		{
			toko.GET("/", tokoController.GetToko)
			toko.PUT("/:id", tokoController.UpdateToko)
		}

		pelangganRepo := repositories.NewMasterPelangganRepository(db)
		pelangganController := controllers.NewMasterPelangganController(pelangganRepo)
		
		adminPelanggan := api.Group("/master-pelanggan")
		adminPelanggan.Use(middleware.AuthMiddleware("Admin"))
		{
			adminPelanggan.GET("/", pelangganController.GetAll)
		}

		pelanggan := api.Group("/master-pelanggan")
		pelanggan.Use(middleware.AuthMiddleware("Admin", "Pelanggan"))
		{
			pelanggan.GET("/:id", pelangganController.GetByID)
			pelanggan.POST("/", pelangganController.Create)
			pelanggan.PUT("/:id", pelangganController.Update)
			pelanggan.DELETE("/:id", pelangganController.Delete)
		}

		alamatPelangganRepo := repositories.NewMasterAlamatPelangganRepository(db)
		alamatPelangganController := controllers.NewMasterAlamatPelangganController(alamatPelangganRepo)
		
		adminAlamatPelanggan := api.Group("/master-alamat-pelanggan")
		adminAlamatPelanggan.Use(middleware.AuthMiddleware("Admin"))
		{
			adminAlamatPelanggan.GET("/", alamatPelangganController.GetAll)
		}

		alamatPelanggan := api.Group("/master-alamat-pelanggan")
		alamatPelanggan.Use(middleware.AuthMiddleware("Admin", "Pelanggan"))
		{
			alamatPelanggan.GET("/:id", alamatPelangganController.GetByID)
			alamatPelanggan.GET("/pelanggan/:id", alamatPelangganController.GetAllByPelanggan)
			alamatPelanggan.POST("/", alamatPelangganController.Create)
			alamatPelanggan.PUT("/:id", alamatPelangganController.Update)
			alamatPelanggan.PUT("/set-alamat-utama/:id", alamatPelangganController.SetAlamatUtama)
			alamatPelanggan.DELETE("/:id", alamatPelangganController.Delete)
		}

		kategoriRepo := repositories.NewMasterKategoriProdukRepository(db)
		kategoriController := controllers.NewMasterKategoriProdukController(kategoriRepo)
		
		kategori := api.Group("/master-kategori-produk")
		kategori.Use(middleware.AuthMiddleware("Admin"))
		{
			kategori.GET("/", kategoriController.GetAll)
			kategori.GET("/:id", kategoriController.GetByID)
			kategori.POST("/", kategoriController.Create)
			kategori.PUT("/:id", kategoriController.Update)
			kategori.DELETE("/:id", kategoriController.Delete)
		}
	}
}
