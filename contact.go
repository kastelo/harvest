package go-harvest

import "fmt"

type ContactService struct {
	Service
}

type Contact struct {
	Id          int         `json:"id"`
	ClientId    int         `json:"client_id"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	Title       string      `json:"title"`
	Email       string      `json:"email"`
	MobilePhone string      `json:"phone_mobile"`
	OfficePhone string      `json:"phone_office"`
	Fax         string      `json:"fax"`
	CreatedAt   HarvestDate `json:"created_at"`
	UpdatedAt   HarvestDate `json:"updated_at"`
}

type ContactResponse struct {
	Contact `json:"contact"`
}

func (c *ContactService) List() (contacts []ContactResponse, err error) {
	resourceURL := "/contacts.json"
	var resp []ContactResponse
	err = c.list(resourceURL, &resp)
	if err != nil {
		return resp, err
	}
	for _, element := range resp {
		contacts = append(contacts, element)
	}
	return contacts, err
}

func (c *ContactService) Find(contactID int) (contact ContactResponse, err error) {
	resourceURL := fmt.Sprintf("/contacts/%v.json", contactID)
	var resp ContactResponse
	err = c.find(resourceURL, &resp)
	return resp, err
}
