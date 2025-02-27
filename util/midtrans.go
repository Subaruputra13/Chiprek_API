package util

import (
	"Chiprek/models"
	"os"

	"github.com/veritrans/go-midtrans"
)

func GetPaymentURL(transaction *models.Transaction, customer *models.Customer) (midtrans.SnapResponse, error) {
	// Set your server key
	midclient := midtrans.NewClient()
	midclient.ServerKey = os.Getenv("Server_Key")
	midclient.ClientKey = os.Getenv("Client_Key")
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	// Create Snap Transaction
	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			FName: customer.Name,
			Phone: customer.PhoneNumber,
		},

		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.TranscationId,
			GrossAmt: int64(transaction.TotalPrice),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return midtrans.SnapResponse{}, err
	}

	return snapTokenResp, nil

}
