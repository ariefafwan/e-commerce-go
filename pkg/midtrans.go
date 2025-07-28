// File: pkg/midtrans.go
package pkg

import (
	"log"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

var SnapClient *snap.Client
var CoreApiClient *coreapi.Client

func InitMidtrans() {
	serverKey := GetEnv("MIDTRANS_SERVER_KEY", "")
	log.Println("Server Key: " + serverKey)
	environment := midtrans.Sandbox
	
	if GetEnv("MIDTRANS_ENVIRONMENT", "sandbox") == "production" {
		environment = midtrans.Production
	}
	
	// ✅ PERBAIKAN TERBAIK: Buat client baru dengan pointer
	SnapClient = &snap.Client{}
	SnapClient.New(serverKey, environment)
	
	CoreApiClient = &coreapi.Client{}
	CoreApiClient.New(serverKey, environment)
	
	log.Println("Midtrans clients initialized successfully")
}

// ✅ Fungsi getter yang aman
func GetSnapClient() *snap.Client {
	if SnapClient == nil {
		log.Println("Warning: SnapClient not initialized, initializing now...")
		InitMidtrans()
	}
	return SnapClient
}