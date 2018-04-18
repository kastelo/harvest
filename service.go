package harvest

import (
	"encoding/json"
)

type Service struct {
	apiClient *APIClient
}

// get makes an HTTP request to the specified path and unmarshals the response
func (service Service) get(resourceURL string, unmarshalContainer interface{}) (err error) {
	contents, err := service.apiClient.GetJSON(resourceURL)
	if err != nil {
		return err
	}
	return json.Unmarshal(contents, unmarshalContainer)
}

func (service Service) post(resourceURL string, v interface{}) (err error) {
	return service.apiClient.PostJSON(resourceURL, v)
}

func (service Service) put(resourceURL string, v interface{}) (err error) {
	return service.apiClient.PutJSON(resourceURL, v)
}
