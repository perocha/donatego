package salesforce_api

import (
	"github.com/simpleforce/simpleforce"
)

// Client is a wrapper around the simpleforce client.
type Client struct {
	sforce *simpleforce.Client
}

// NewClient creates a new client.
func NewClient(instanceUrl string, clientId string, version string) (*Client, error) {
	sforce := simpleforce.NewClient(instanceUrl, clientId, version)
	return &Client{sforce}, nil
}

// Authenticate method deals with the authentication with Salesforce.
func (c *Client) Authenticate(username string, password string, securityToken string) error {
	err := c.sforce.LoginPassword(username, password, securityToken)
	if err != nil {
		return err
	}
	return nil
}

// Query method deals with the query to Salesforce.
func (c *Client) Query(query string) ([]map[string]interface{}, error) {
	res, err := c.sforce.Query(query)
	if err != nil {
		return nil, err
	}

	var records []map[string]interface{}
	for _, obj := range res.Records {
		record := make(map[string]interface{})
		for key := range obj {
			val := obj[key]
			record[key] = val
		}
		records = append(records, record)
	}

	return records, nil
}
