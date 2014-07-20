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

func (c *ClientService) List() (err error, clients []ClientResponse) {
	resourceURL := "/clients.json"
	var resp []ClientResponse
	err = c.list(resourceURL, &resp)
	if err != nil {
		return err, resp
	}
	for _, element := range resp {
		clients = append(clients, element)
	}
	return err, clients
}

func (c *ClientService) Find(clientID int) (err error, client ClientResponse) {
	resourceURL := fmt.Sprintf("/clients/%v.json", clientID)
	var resp ClientResponse
	err = c.find(resourceURL, &resp)
	return err, resp
}
