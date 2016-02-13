package harvest

import (
	"encoding/json"
)

type Service struct {
	apiClient *APIClient
}

// find makes an HTTP request to the specified path and returns
// the unmarshalled HTTP response in
func (service Service) find(resourceURL string, unmarshalContainer interface{}) (err error) {
	contents, err := service.apiClient.GetJSON(resourceURL)
	if err != nil {
		return err
	}
	return json.Unmarshal(contents, unmarshalContainer)
}

// list makes an HTTP request to the specified path and returns
// the address of the unmarshalled HTTP response
func (service Service) list(resourceURL string, unmarshalContainer interface{}) (err error) {
	contents, err := service.apiClient.GetJSON(resourceURL)
	if err != nil {
		return err
	}
	return json.Unmarshal(contents, &unmarshalContainer)
}
