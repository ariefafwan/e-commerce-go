package main

import (
	"e-commerce-go/pkg"
	"e-commerce-go/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	pkg.ConnectDB()
	pkg.InitCloudinary()

	r := gin.Default()
	routers.SetupRouters(r, pkg.DB)

	r.Run(":" + pkg.GetEnv("APP_PORT", "8080"))
	fmt.Println("Server berjalan di port:" + pkg.GetEnv("APP_PORT", "8080"))
}