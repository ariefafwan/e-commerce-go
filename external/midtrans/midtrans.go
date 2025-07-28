package midtrans

import (
	"e-commerce-go/internal/dto"
	"e-commerce-go/pkg"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func getSnapClient() *snap.Client {
	return pkg.GetSnapClient()
}
// var recoreapiclient = pkg.CoreApiClient

// dokumentasi ada di : https://docs.midtrans.com/docs/snap-snap-integration-guide
func CreatePayment(transaksi *dto.TransaksiResponse) (payment_token string, payment_url string, id_transaksi string, error error) {
	// Format item details
	var itemDetails []midtrans.ItemDetails
	for _, item := range transaksi.DataItems {
		itemDetails = append(itemDetails, midtrans.ItemDetails{
			ID:    item.IDProduk.String(),
			Name:  item.DataProduk.Nama,
			Price: int64(item.DataVariant.Harga),
			Qty:   int32(item.Quantity),
		})
	}

	// tambahin ongkir sebagai item
	if transaksi.TotalOngkir > 0 {
		itemDetails = append(itemDetails, midtrans.ItemDetails{
			ID:    "SHIPPING",
			Name:  "Biaya Pengiriman",
			Price: int64(transaksi.TotalOngkir),
			Qty:   1,
		})
	}

	// tambahin pajak sebagai item
	if transaksi.Pajak > 0 {
		itemDetails = append(itemDetails, midtrans.ItemDetails{
			ID:    "TAX",
			Name:  "Pajak",
			Price: int64(transaksi.Pajak),
			Qty:   1,
		})
	}

	// detail pelanggan
	customerDetails := &midtrans.CustomerDetails{
		FName: transaksi.DataPelanggan.NamaLengkap,
		Email: transaksi.DataPelanggan.DataUser.Email,
		Phone: transaksi.DataPelanggan.Phone,
		BillAddr: &midtrans.CustomerAddress{
			FName:       transaksi.DataPelanggan.NamaLengkap,
			Phone:       transaksi.DataPelanggan.Phone,
			Address:     transaksi.DataAlamat.Label,
			City:        transaksi.DataAlamat.DataKecamatan.DataKota.Nama,
			Postcode:    transaksi.DataAlamat.KodePos,
			CountryCode: "IDN",
		},
		ShipAddr: &midtrans.CustomerAddress{
			FName:       transaksi.DataPelanggan.NamaLengkap,
			Phone:       transaksi.DataPelanggan.Phone,
			Address:     transaksi.DataAlamat.Label,
			City:        transaksi.DataAlamat.DataKecamatan.DataKota.Nama,
			Postcode:    transaksi.DataAlamat.KodePos,
			CountryCode: "IDN",
		},
	}

	expiryDuration := int64(time.Until(*transaksi.ExpiredAt).Hours())

	// set snap request
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  *transaksi.NoInvoice,
			GrossAmt: int64(transaksi.GrandTotal),
		},
		Items:           &itemDetails,
		CustomerDetail:  customerDetails,
		EnabledPayments: snap.AllSnapPaymentType,
		Expiry: &snap.ExpiryDetails{
			StartTime: time.Now().Format("2006-01-02 15:04:05 -0700"),
			Unit:      "hours",
			Duration:  expiryDuration,
		},
		Callbacks: &snap.Callbacks{
			Finish: pkg.GetEnv("MIDTRANS_CALLBACK_URL", ""),
		},
	}

	// bukan transaksi, untuk ambil token dan url nya
	snapClient := getSnapClient()
	snapResp, err := snapClient.CreateTransaction(req)
	if err != nil {
		return "", "", "", err
	}

	return snapResp.Token, snapResp.RedirectURL, *transaksi.NoInvoice, nil
}

func HandleCallback(orderID string, transactionStatus string, fraudStatus string) (string, error) {
	// Mapping status midtrans ke status aplikasi
	var status string
	
	switch transactionStatus {
	case "capture":
		if fraudStatus == "challenge" {
			status = "Pending"
		} else if fraudStatus == "accept" {
			status = "Paid"
		}
	case "settlement":
		status = "Paid"
	case "cancel", "deny":
		status = "Cancelled"
	case "expire":
		status = "Expired"
	case "pending":
		status = "Pending"
	default:
		status = "Pending"
	}

	return status, nil
}