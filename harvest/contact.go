package harvest

import (
	"fmt"
)

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
	Contact Contact
}

func (c *ContactService) List() (err error, contacts []Contact) {
	resourceURL := "/contacts.json"
	var contactResponse []ContactResponse
	err = c.list(resourceURL, &contactResponse)
	if err != nil {
		return
	}

	for _, element := range contactResponse {
		contacts = append(contacts, element.Contact)
	}
	return
}

func (c *ContactService) Find(contactID int) (err error, contact Contact) {
	resourceURL := fmt.Sprintf("/contacts/%v.json", contactID)
	var contactResponse ContactResponse
	err = c.find(resourceURL, &contactResponse)
	if err != nil {
		return
	}
	contact = contactResponse.Contact
	return
}
