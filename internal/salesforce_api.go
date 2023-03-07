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

/*
func (c *Client) Query(query string) ([]map[string]interface{}, error) {
	res, err := c.sforce.Query(query)
	if err != nil {
		return nil, err
	}

	var records []map[string]interface{}
	for _, obj := range res.Records {
		record := make(map[string]interface{})
		for key, val := range obj.MapData {
			record[key] = val
		}
		records = append(records, record)
	}

	return records, nil
}
*/
/*
func main() {
	// Set the configuration file name.
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Get the configuration from the config file.
	clientID := viper.GetString("clientID")
	username := viper.GetString("username")
	password := viper.GetString("password")
	securityToken := viper.GetString("securityToken")

	// Create a new client.
	client := simpleforce.NewClient("https://login.salesforce.com", clientID, "52.0")

	// Login using username and password.
	sf_err := client.LoginPassword(username, password, securityToken)
	if sf_err != nil {
		fmt.Println(err.Error())
		return
	}

	// Query accounts.
	qResult, err := client.Query("SELECT Id, Name FROM Account")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, record := range qResult.Records {
		fmt.Printf("%s: %s\n", record["Id"], record["Name"])
	}
}
*/
