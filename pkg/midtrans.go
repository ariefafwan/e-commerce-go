package pkg

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

var SnapClient *snap.Client
var CoreApiClient *coreapi.Client

func InitMidtrans() {
	var snapClient snap.Client
	var coreApiClient coreapi.Client
	
	serverKey := GetEnv("MIDTRANS_SERVER_KEY", "")
	environment := midtrans.Sandbox
	
	// production kapan kapan
	if GetEnv("APP_ENV", "development") == "production" {
		environment = midtrans.Production
	}
	
	snapClient.New(serverKey, environment)
	coreApiClient.New(serverKey, environment)
}