package payment

import (
	"fmt"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransService struct {
	snapClient    snap.Client
	coreApiClient coreapi.Client
	environment   midtrans.EnvironmentType
}

func NewMidtransService(serverKey string, isProduction bool) *MidtransService {
	s := &MidtransService{}

	if isProduction {
		s.environment = midtrans.Production
	} else {
		s.environment = midtrans.Sandbox
	}

	s.snapClient.New(serverKey, s.environment)
	s.coreApiClient.New(serverKey, s.environment)

	return s
}

func (m *MidtransService) CreateTransaction(orderID string, amount int64, customerName, customerEmail, itemName string) (*snap.Response, error) {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: amount,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: customerName,
			Email: customerEmail,
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    orderID,
				Name:  itemName,
				Price: amount,
				Qty:   1,
			},
		},
		EnabledPayments: snap.AllSnapPaymentType,
	}

	resp, err := m.snapClient.CreateTransaction(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %v", err)
	}

	return resp, nil
}

// GetTransactionDetails dapatkan transaction_id dari Core API
func (m *MidtransService) GetTransactionDetails(orderID string) (*coreapi.TransactionStatusResponse, error) {
	resp, err := m.coreApiClient.CheckTransaction(orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction details: %v", err)
	}
	return resp, nil
}

func (m *MidtransService) CheckTransactionStatus(orderID string) (*coreapi.TransactionStatusResponse, error) {
	return m.GetTransactionDetails(orderID)
}
