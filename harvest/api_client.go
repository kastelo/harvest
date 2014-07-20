package harvest

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"code.google.com/p/goauth2/oauth"
)

type APIClient struct {
	username   string
	password   string
	subdomain  string
	httpClient *http.Client

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

func NewAPIClientWithBasicAuth(username, password, subdomain string) (c *APIClient) {
	c = newAPIClient(subdomain, nil)
	c.username = username
	c.password = password
	return c
}

func NewAPIClientWithAuthToken(token, subdomain string) (c *APIClient) {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: token},
	}
	c = newAPIClient(subdomain, t.Client())
	return c
}

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
