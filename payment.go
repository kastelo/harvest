package harvest

import (
	"fmt"
)

type PaymentService struct {
	Service
}

type Payment struct {
	ID                  int     `json:"id"`
	InvoiceID           int     `json:"invoice_id"`
	Amount              float64 `json:"amount"`
	PaidAt              Date    `json:"paid_at"`
	CreatedAt           Date    `json:"created_at"`
	Notes               string  `json:"notes"`
	RecordedBy          string  `json:"recorded_by"`
	RecordedByEmail     string  `json:"recorded_by_email"`
	PayPalTransactionID *int    `json:"pay_pal_transaction_id"`
	Authorization       *int    `json:"authorization"`
	PaymentGatewayID    *int    `json:"payment_gateway_id"`
}

type PaymentResponse struct {
	Payment `json:"payment"`
}

// List requests list of payments and for an invoice
func (p *PaymentService) ListInvoice(invoiceID int) ([]PaymentResponse, error) {
	resourceURL := fmt.Sprintf("/invoices/%d/payments.json", invoiceID)
	var payments []PaymentResponse
	if err := p.get(resourceURL, &payments); err != nil {
		return nil, err
	}
	return payments, nil
}
