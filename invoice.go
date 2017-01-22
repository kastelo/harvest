package harvest

import (
	"fmt"
	"time"
)

type InvoiceService struct {
	Service
}

type Invoice struct {
	ID                 int       `json:"id"`
	Amount             float32   `json:"amount"`
	DueAmount          float32   `json:"due_amount"`
	DueAt              Date      `json:"due_at"`
	DueAtHumanFormat   string    `json:"due_at_human_format"`
	PeriodEnd          Date      `json:"period_end"`
	PeriodStart        Date      `json:"period_start"`
	ClientID           int       `json:"client_id"`
	Subject            string    `json:"subject"`
	Currency           string    `json:"currency"`
	IssuedAt           Date      `json:"issued_at"`
	CreatedByID        int       `json:"created_by_id"`
	Notes              string    `json:"notes"`
	Number             string    `json:"number"`
	PurchaseOrder      string    `json:"purchase_order"`
	ClientKey          string    `json:"client_key"`
	State              string    `json:"state"`
	Tax                float32   `json:"tax"`
	Tax2               float32   `json:"tax2"`
	TaxAmount          float32   `json:"tax_amount"`
	TaxAmount2         float32   `json:"tax2_amount"`
	DiscountAmount     float32   `json:"discount_amount"`
	Discount           float32   `json:"discount"`
	RecurringInvoiceID int       `json:"recurring_invoice_id"`
	EstimateID         int       `json:"estimate_id"`
	RetainerID         int       `json:"retainer_id"`
	UpdatedAt          time.Time `json:"updated_at"`
	CreatedAt          time.Time `json:"created_at"`
}

type InvoiceResponse struct {
	Invoice `json:"invoices"`
}

// List requests list of invoices and returns response
func (i *InvoiceService) List() (invoices []InvoiceResponse, err error) {
	resourceURL := "/invoices.json"
	var resp []InvoiceResponse
	err = i.list(resourceURL, &resp)
	if err != nil {
		return resp, err
	}
	for _, element := range resp {
		invoices = append(invoices, element)
	}
	return invoices, err
}

// Find requests invoice information for specified invoice and returns response
func (i *InvoiceService) Find(invoiceID int) (invoice InvoiceResponse, err error) {
	resourceURL := fmt.Sprintf("/invoices/%v.json", invoiceID)
	var resp InvoiceResponse
	err = i.find(resourceURL, &resp)
	return resp, err
}
