package main

import (
	"fmt"

	"github.com/perocha/internal/salesforce_api"
	"github.com/spf13/viper"
)

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

	fmt.Println(clientID)
	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(securityToken)

	// Create a new Salesforce client
	sf := salesforce_api.NewSalesforceClient("https://login.salesforce.com", clientID, "52.0")

	// Authenticate with Salesforce
	sf_err := sf.Authenticate(username, password, securityToken)
	if sf_err != nil {
		panic(sf_err)
	}
	/*
		// Query Salesforce
		records, err := sf.Query("SELECT Id, Name FROM Contact")
		if err != nil {
			panic(err)
		}
	*/
}
