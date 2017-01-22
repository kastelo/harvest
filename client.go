package harvest

import (
	"fmt"
	"time"
)

type ClientService struct {
	Service
}

type Client struct {
	Name                    string    `json:"name"`
	Currency                string    `json:"currency"`
	Active                  bool      `json:"active"`
	ID                      int       `json:"id"`
	HighriseID              int       `json:"highrise_id"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	Details                 string    `json:"details"`
	DefaultInvoiceTimeframe string    `json:"default_invoice_timeframe"`
	LastInvoiceKind         string    `json:"last_invoice_kind"`
}

type ClientResponse struct {
	Client `json:"client"`
}

// List requests list of clients and returns response
func (c *ClientService) List() ([]ClientResponse, error) {
	resourceURL := "/clients.json"
	var clients []ClientResponse
	if err := c.get(resourceURL, &clients); err != nil {
		return nil, err
	}
	return clients, nil
}

// Find requests client information for specified client and returns response
func (c *ClientService) Find(clientID int) (ClientResponse, error) {
	resourceURL := fmt.Sprintf("/clients/%v.json", clientID)
	var resp ClientResponse
	err := c.get(resourceURL, &resp)
	return resp, err
}
