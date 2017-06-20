// +build oauth2

package harvest

import (
	"context"

	"golang.org/x/oauth2"
)

// NewAPIClientWithAuthToken instantiates a new http.Client and returns a new
// APIClient using an OAuth token
func NewAPIClientWithAuthToken(token, subdomain string) (c *APIClient) {
	client := oauth2.NewClient(
		context.Background(),
		oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token}))
	c = newAPIClient(subdomain, client, 0)
	return c
}
