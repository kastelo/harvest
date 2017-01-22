package harvest

import "fmt"

type ContactService struct {
	Service
}

type Contact struct {
	ID          int    `json:"id"`
	ClientID    int    `json:"client_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Title       string `json:"title"`
	Email       string `json:"email"`
	MobilePhone string `json:"phone_mobile"`
	OfficePhone string `json:"phone_office"`
	Fax         string `json:"fax"`
	CreatedAt   Date   `json:"created_at"`
	UpdatedAt   Date   `json:"updated_at"`
}

type ContactResponse struct {
	Contact `json:"contact"`
}

// List requests list of contacts and returns response
func (c *ContactService) List() ([]ContactResponse, error) {
	resourceURL := "/contacts.json"
	var contacts []ContactResponse
	if err := c.get(resourceURL, &contacts); err != nil {
		return nil, err
	}
	return contacts, nil
}

// Find requests contact information for specified contact and returns repsonse
func (c *ContactService) Find(contactID int) (ContactResponse, error) {
	resourceURL := fmt.Sprintf("/contacts/%v.json", contactID)
	var resp ContactResponse
	err := c.get(resourceURL, &resp)
	return resp, err
}
