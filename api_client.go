/*
Package harvest provides data structures and a wrapper for the Harvest API
*/
package harvest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

// APIClient contains credentials & data interfaces
type APIClient struct {
	// Authentication and connection information
	username   string
	password   string
	subdomain  string
	httpClient *http.Client

	// caching
	cacheGroup    *singleflight.Group
	cacheInterval time.Duration
	cacheEntries  map[string]cacheEntry // url -> response data
	cacheMut      sync.Mutex

	// Data interface accessors
	Client          *ClientService
	People          *PersonService
	Project         *ProjectService
	Invoice         *InvoiceService
	Account         *AccountService
	Task            *TaskService
	Contact         *ContactService
	Expense         *ExpenseService
	ExpenseCategory *ExpenseCategoryService
	Entry           *EntryService
	Payment         *PaymentService
}

type cacheEntry struct {
	when time.Time
	data []byte
}

// newAPIClient instantiates a new http.Client and returns a new APIClient
func newAPIClient(subdomain string, httpClient *http.Client, cacheInterval time.Duration) (c *APIClient) {
	c = new(APIClient)
	c.subdomain = subdomain
	c.cacheGroup = new(singleflight.Group)
	c.cacheInterval = cacheInterval
	c.cacheEntries = make(map[string]cacheEntry)

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
	c.Expense = &ExpenseService{Service{c}}
	c.ExpenseCategory = &ExpenseCategoryService{Service{c}}
	c.Entry = &EntryService{Service{c}}
	c.Payment = &PaymentService{Service{c}}
	return c
}

// NewAPIClientWithBasicAuth instantiates a new http.Client and returns a new
// APIClient using HTTP basic auth credentials
func NewAPIClientWithBasicAuth(username, password, subdomain string) (c *APIClient) {
	return NewCachingAPIClientWithBasicAuth(username, password, subdomain, 0)
}

// NewCachingAPIClientWithBasicAuth instantiates a new http.Client and
// returns a new APIClient using HTTP basic auth credentials. Responses are
// cached and reused for up to cacheInterval.
func NewCachingAPIClientWithBasicAuth(username, password, subdomain string, cacheInterval time.Duration) (c *APIClient) {
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

	c = newAPIClient(subdomain, nil, cacheInterval)
	c.username = username
	c.password = password
	return c
}

// GetJSON makes an HTTP GET request to the specified path and returns the body
// of the HTTP response
func (c *APIClient) GetJSON(path string) (jsonResponse []byte, err error) {
	v, err, _ := c.cacheGroup.Do(path, func() (interface{}, error) {
		c.cacheMut.Lock()
		cur := c.cacheEntries[path]
		c.cacheMut.Unlock()

		if time.Since(cur.when) < c.cacheInterval {
			return cur.data, nil
		}

		data, err := c.uncachedGetJSON(path)
		if err != nil {
			return nil, err
		}

		c.cacheMut.Lock()
		c.cacheEntries[path] = cacheEntry{
			when: time.Now(),
			data: data,
		}
		c.cacheMut.Unlock()

		return data, nil
	})
	if err != nil {
		return nil, err
	}
	return v.([]byte), nil
}

func (c *APIClient) uncachedGetJSON(path string) (jsonResponse []byte, err error) {
	resourceURL := fmt.Sprintf("https://%v.harvestapp.com%v", c.subdomain, path)
	request, err := http.NewRequest(http.MethodGet, resourceURL, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(c.username, c.password)
	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return bs, err
	}

	if resp.StatusCode != http.StatusOK {
		return bs, errors.New(resp.Status)
	}

	return bs, nil
}

func (c *APIClient) PostJSON(path string, v interface{}) error {
	bs, err := json.Marshal(v)
	if err != nil {
		return err
	}

	r := bytes.NewReader(bs)
	resourceURL := fmt.Sprintf("https://%v.harvestapp.com%v", c.subdomain, path)
	request, err := http.NewRequest(http.MethodPost, resourceURL, r)
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(c.username, c.password)
	resp, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}

	ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return errors.New(resp.Status)
	}

	return nil
}
