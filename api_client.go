/*
Package harvest provides data structures and a wrapper for the Harvest API
*/
package go-harvest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"golang.org/x/net/context"
	"os"
	"strings"

	"golang.org/x/oauth2"
)

// APIClient contains credentials & data interfaces
type APIClient struct {
	// Authentication and connection information
	username   string
	password   string
	subdomain  string
	httpClient *http.Client

	// Data interface accessors
	Client          *ClientService
	People          *PersonService
	Project         *ProjectService
	Invoice         *InvoiceService
	Account         *AccountService
	Task            *TaskService
	Contact         *ContactService
	ExpenseCategory *ExpenseCategoryService
	Entry           *EntryService
}

func newAPIClient(subdomain string, httpClient *http.Client) (c *APIClient) {
	c = new(APIClient)
	c.subdomain = subdomain

	if httpClient != nil {
		c.httpClient = httpClient
	} else {
		c.httpClient = new(http.Client)
	}

	c.Client = &ClientService{Service{c}}
	c.People = &PersonService{Service{c}}
	c.Project = &ProjectService{Service{c}}
	c.Invoice = &InvoiceService{Service{c}}
	c.Account = &AccountService{Service{c}}
	c.Task = &TaskService{Service{c}}
	c.Contact = &ContactService{Service{c}}
	c.ExpenseCategory = &ExpenseCategoryService{Service{c}}
	c.Entry = &EntryService{Service{c}}
	return c
}

// NewAPIClientWithBasicAuth instantiates a new http.Client and returns a new
// APIClient using HTTP basic auth credentials
func NewAPIClientWithBasicAuth(username, password, subdomain string) (c *APIClient) {
	var missingData []string
	if subdomain == "" {
		missingData = append(missingData, "subdomain")
	}
	if username == "" {
		missingData = append(missingData, "username")
	}
	if password == "" {
		missingData = append(missingData, "password")
	}
	errorMsg := strings.Join(missingData, ", ")
	if errorMsg != "" {
		fmt.Println("ERROR! You are missing the following:", errorMsg)
		os.Exit(1)
	}

	c = newAPIClient(subdomain, nil)
	c.username = username
	c.password = password
	return c
}

// NewAPIClientWithAuthToken instantiates a new http.Client and returns a new
// APIClient using an OAuth token
func NewAPIClientWithAuthToken(token, subdomain string) (c *APIClient) {
	client := oauth2.NewClient(
		context.Background(),
		oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token}))
	c = newAPIClient(subdomain, client)
	return c
}

// GetJSON makes an HTTP GET request to the specified path and returns the body
// of the HTTP response
func (c *APIClient) GetJSON(path string) (jsonResponse []byte, err error) {
	resourceURL := fmt.Sprintf("https://%v.harvestapp.com%v", c.subdomain, path)
	request, err := http.NewRequest("GET", resourceURL, nil)
	if err != nil {
		return []byte(""), err
	}

	request.SetBasicAuth(c.username, c.password)
	resp, err := c.httpClient.Do(request)

	if err != nil {
		return []byte(""), err
	}

	defer resp.Body.Close()
	jsonResponse, err = ioutil.ReadAll(resp.Body)
	return jsonResponse, err
}
