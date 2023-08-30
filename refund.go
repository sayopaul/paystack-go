package paystack

import (
	"time"
)

// RefundService handles operations related to transactions
// For more details see https://developers.paystack.co/v1.0/reference#create-transaction
type RefundService service

// RefundRequest represents a request to start a Refund.
type RefundRequest struct {
	Transaction  string  `json:"transaction,omitempty"`
	Currency     string  `json:"currency,omitempty"`
	Amount       float32 `json:"amount,omitempty"`
	MerchantNote string  `json:"merchant_note,omitempty"`
	CustomerNote string  `json:"customer_note,omitempty"`
}

type RefundResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Transaction struct {
			ID            int       `json:"id"`
			Domain        string    `json:"domain"`
			Reference     string    `json:"reference"`
			Amount        int       `json:"amount"`
			PaidAt        time.Time `json:"paid_at"`
			Channel       string    `json:"channel"`
			Currency      string    `json:"currency"`
			Authorization struct {
				ExpMonth    interface{} `json:"exp_month"`
				ExpYear     interface{} `json:"exp_year"`
				AccountName interface{} `json:"account_name"`
			} `json:"authorization"`
			Customer struct {
				InternationalFormatPhone interface{} `json:"international_format_phone"`
			} `json:"customer"`
			Plan struct {
			} `json:"plan"`
			Subaccount struct {
				Currency interface{} `json:"currency"`
			} `json:"subaccount"`
			Split struct {
			} `json:"split"`
			OrderID            interface{} `json:"order_id"`
			PaidAtt            time.Time   `json:"paidAt"`
			PosTransactionData interface{} `json:"pos_transaction_data"`
			Source             interface{} `json:"source"`
			FeesBreakdown      interface{} `json:"fees_breakdown"`
		} `json:"transaction"`
		Integration    int         `json:"integration"`
		DeductedAmount int         `json:"deducted_amount"`
		Channel        interface{} `json:"channel"`
		MerchantNote   string      `json:"merchant_note"`
		CustomerNote   string      `json:"customer_note"`
		Status         string      `json:"status"`
		RefundedBy     string      `json:"refunded_by"`
		ExpectedAt     time.Time   `json:"expected_at"`
		Currency       string      `json:"currency"`
		Domain         string      `json:"domain"`
		Amount         int         `json:"amount"`
		FullyDeducted  bool        `json:"fully_deducted"`
		ID             int         `json:"id"`
		CreatedAt      time.Time   `json:"createdAt"`
		UpdatedAt      time.Time   `json:"updatedAt"`
	} `json:"data"`
}

// Refund is for refunding payments.
// For more details see https://paystack.com/docs/api/refund/
func (s *RefundService) Refund(req *RefundRequest) (*RefundResponse, error) {
	rf := &RefundResponse{}
	err := s.client.CallSimple("POST", "/refund", req, rf)
	return rf, err
}
