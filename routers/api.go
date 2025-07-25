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

	rajaOngkirRepo := repositories.NewRajaOngkirRepository(db)
	rajaOngkirController := controllers.NewRajaOngkirController(rajaOngkirRepo)
	
	router.GET("/provinsi", rajaOngkirController.GetProvinsi)
	router.GET("/kota/:id", rajaOngkirController.GetKota)
	router.GET("/kecamatan/:id", rajaOngkirController.GetKecamatan)

	api := router.Group("/api")
	{
		rajaOngkirRepo := repositories.NewRajaOngkirRepository(db)
		rajaOngkirController := controllers.NewRajaOngkirController(rajaOngkirRepo)
		
		api.GET("/provinsi", rajaOngkirController.GetProvinsi)
		api.GET("/kota/:id", rajaOngkirController.GetKota)
		api.GET("/kecamatan/:id", rajaOngkirController.GetKecamatan)

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

		produkRepo := repositories.NewMasterProdukRepository(db)
		produkVariantRepo := repositories.NewMasterProdukVariantRepository(db)
		produkGaleriRepo := repositories.NewMasterProdukGaleriRepository(db)
		produkController := controllers.NewMasterProdukController(produkRepo, produkVariantRepo, produkGaleriRepo)

		AdminProduk := api.Group("/master-produk")
		AdminProduk.Use(middleware.AuthMiddleware("Admin"))
		{
			AdminProduk.POST("/", produkController.Create)
			AdminProduk.GET("/non-aktif", produkController.GetProdukNonAktif)
			AdminProduk.PUT("/:slug", produkController.Update)
			AdminProduk.DELETE("/:slug", produkController.Delete)

			produkVariant := AdminProduk.Group("/:slug/variant")
			{
				produkVariant.GET("/", produkController.GetAllVariant)
				produkVariant.GET("/:id", produkController.GetVariantByID)
				produkVariant.POST("", produkController.CreateVariant)
				produkVariant.PUT("/:id", produkController.UpdateVariant)
				produkVariant.DELETE("/:id", produkController.DeleteVariant)
			}

			produkGaleri := AdminProduk.Group("/:slug/galeri")
			{
				produkGaleri.GET("/", produkController.GetAllGaleri)
				produkGaleri.GET("/:id", produkController.GetGaleriByID)
				produkGaleri.POST("", produkController.CreateGaleri)
				produkGaleri.PUT("/:id", produkController.UpdateGaleri)
				produkGaleri.DELETE("/:id", produkController.DeleteGaleri)
			}
		}

		produk := api.Group("/master-produk")
		produk.Use(middleware.AuthMiddleware("Admin", "Pelanggan"))
		{
			produk.GET("/", produkController.GetAll)
			produk.GET("/kategori/:slug", produkController.GetAllByKategori)
			produk.GET("/:slug", produkController.GetBySlug)
		}

		transaksiKeranjangItemRepo := repositories.NewTransaksiKeranjangItemRepository(db)
		transaksiKeranjangRepo := repositories.NewTransaksiKeranjangRepository(db)
		keranjangController := controllers.NewTransaksiKeranjangController(transaksiKeranjangRepo, transaksiKeranjangItemRepo, produkVariantRepo)

		AdminKeranjang := api.Group("/keranjang")
		AdminKeranjang.Use(middleware.AuthMiddleware("Admin"))
		{
			AdminKeranjang.GET("/", keranjangController.GetAllKeranjang)
		}

		keranjang := api.Group("/keranjang")
		keranjang.Use(middleware.AuthMiddleware("Admin"," Pelanggan"))
		{
			keranjang.GET("/pelanggan/:id", keranjangController.GetAllByPelanggan)
			keranjang.GET("/:id", keranjangController.GetByID)
			keranjang.POST("/", keranjangController.Create)
			keranjang.PUT("/:id", keranjangController.Update)
			keranjang.PUT("/item/:id", keranjangController.UpdateItem)
			keranjang.DELETE("/:id", keranjangController.Delete)
			keranjang.DELETE("/item/:id", keranjangController.DeleteItem)
		}
	}
}
