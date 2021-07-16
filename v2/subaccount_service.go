package binance

import (
	"context"
	"encoding/json"
)

// CreateOrderService create order
type TransferToSubAccountService struct {
	c       *Client
	toEmail string
	asset   string
	amount  *string
}

// ToEmail set toEmail
func (s *TransferToSubAccountService) ToEmail(toEmail string) *TransferToSubAccountService {
	s.toEmail = toEmail
	return s
}

// Asset set asset
func (s *TransferToSubAccountService) Asset(asset string) *TransferToSubAccountService {
	s.asset = asset
	return s
}

// Amount set amount
func (s *TransferToSubAccountService) Amount(amount string) *TransferToSubAccountService {
	s.amount = &amount
	return s
}

func (s *TransferToSubAccountService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"toEmail": s.toEmail,
		"asset":   s.asset,
	}
	if s.amount != nil {
		m["amount"] = *s.amount
	}
	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *TransferToSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *CreateOrderResponse, err error) {
	data, err := s.createOrder(ctx, "/sapi/v1/sub-account/transfer/subToSub", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Test send test api to check if the request is valid
func (s *TransferToSubAccountService) Test(ctx context.Context, opts ...RequestOption) (err error) {
	_, err = s.createOrder(ctx, "/sapi/v1/sub-account/transfer/subToSub/test", opts...)
	return err
}

// CreateOrderResponse define create order response
type TransferToSubAccountResponse struct {
	TxnID string `json:"txnId"`
}
