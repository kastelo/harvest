package harvest

import (
	"fmt"
	"time"
)

type InvoiceService struct {
	Service
}

type Invoice struct {
	Id                 int         `json:"id"`
	Amount             float32     `json:"amount"`
	DueAmount          float32     `json:"due_amount"`
	DueAt              HarvestDate `json:"due_at"`
	DueAtHumanFormat   string      `json:"due_at_human_format"`
	PeriodEnd          HarvestDate `json:"period_end"`
	PeriodStart        HarvestDate `json:"period_start"`
	ClientId           int         `json:"client_id"`
	Subject            string      `json:"subject"`
	Currency           string      `json:"currency"`
	IssuedAt           HarvestDate `json:"issued_at"`
	CreatedById        int         `json:"created_by_id"`
	Notes              string      `json:"notes"`
	Number             string      `json:"number"`
	PurchaseOrder      string      `json:"purchase_order"`
	ClientKey          string      `json:"client_key"`
	State              string      `json:"state"`
	Tax                float32     `json:"tax"`
	Tax2               float32     `json:"tax2"`
	TaxAmount          float32     `json:"tax_amount"`
	TaxAmount2         float32     `json:"tax_amount2"`
	DiscountAmount     float32     `json:"discount_amount"`
	Discount           float32     `json:"discount"`
	RecurringInvoiceId int         `json:"recurring_invoice_id"`
	EstimateId         int         `json:"estimate_id"`
	RetainerId         int         `json:"retainer_id"`
	UpdatedAt          time.Time   `json:"updated_at"`
	CreatedAt          time.Time   `json:"created_at"`
}

type InvoiceResponse struct {
	Invoice `json:"invoices"`
}

func (i *InvoiceService) List() (err error, invoices []InvoiceResponse) {
	resourceURL := "/invoices.json"
	var resp []InvoiceResponse
	err = i.list(resourceURL, &resp)
	if err != nil {
		return err, resp
	}
	for _, element := range resp {
		invoices = append(invoices, element)
	}
	return err, invoices
}

func (i *InvoiceService) Find(invoiceID int) (err error, invoice InvoiceResponse) {
	resourceURL := fmt.Sprintf("/invoices/%v.json", invoiceID)
	var resp InvoiceResponse
	err = i.find(resourceURL, &resp)
	return err, resp
}
