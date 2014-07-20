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
	Id                      int       `json:"id"`
	HighriseId              int       `json:"highrise_id"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"created_at"`
	Details                 string    `json:"details"`
	DefaultInvoiceTimeframe string    `json:"default_invoice_timeframe"`
	LastInvoiceKind         string    `json:"last_invoice_kind"`
}

type ClientResponse struct {
	Client `json:"client"`
}

func (c *ClientService) List() (clients []ClientResponse, err error) {
	resourceURL := "/clients.json"
	var resp []ClientResponse
	err = c.list(resourceURL, &resp)
	if err != nil {
		return resp, err
	}
	for _, element := range resp {
		clients = append(clients, element)
	}
	return clients, err
}

func (c *ClientService) Find(clientID int) (client ClientResponse, err error) {
	resourceURL := fmt.Sprintf("/clients/%v.json", clientID)
	var resp ClientResponse
	err = c.find(resourceURL, &resp)
	return resp, err
}
