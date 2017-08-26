package harvest

import (
	"encoding/csv"
	"fmt"
	"strings"
	"time"
)

type InvoiceService struct {
	Service
}

type Invoice struct {
	ID                 int       `json:"id"`
	Amount             float64   `json:"amount"`
	DueAmount          float64   `json:"due_amount"`
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
	Tax                float64   `json:"tax"`
	Tax2               float64   `json:"tax2"`
	TaxAmount          float64   `json:"tax_amount"`
	TaxAmount2         float64   `json:"tax2_amount"`
	DiscountAmount     float64   `json:"discount_amount"`
	Discount           float64   `json:"discount"`
	RecurringInvoiceID int       `json:"recurring_invoice_id"`
	EstimateID         int       `json:"estimate_id"`
	RetainerID         int       `json:"retainer_id"`
	UpdatedAt          time.Time `json:"updated_at"`
	CreatedAt          time.Time `json:"created_at"`
	CSVLineItems       string    `json:"csv_line_items"`
}

type InvoiceResponse struct {
	Invoice `json:"invoices"`
}

// List requests list of invoices and returns response
func (i *InvoiceService) List() ([]InvoiceResponse, error) {
	resourceURL := "/invoices.json"
	var invoices []InvoiceResponse
	if err := i.get(resourceURL, &invoices); err != nil {
		return nil, err
	}
	return invoices, nil
}

// Find requests invoice information for specified invoice and returns response
func (i *InvoiceService) Find(invoiceID int) (InvoiceResponse, error) {
	resourceURL := fmt.Sprintf("/invoices/%v.json", invoiceID)
	var resp InvoiceResponse
	err := i.get(resourceURL, &resp)
	return resp, err
}

func (i *Invoice) LineItems() []map[string]string {
	cr := csv.NewReader(strings.NewReader(i.CSVLineItems))
	lines, err := cr.ReadAll()
	if err != nil || len(lines) < 2 {
		return nil
	}

	fieldNames := lines[0]
	var res []map[string]string
	for _, line := range lines[1:] {
		m := make(map[string]string)
		for i, k := range fieldNames {
			m[k] = line[i]
		}
		res = append(res, m)
	}
	return res
}
