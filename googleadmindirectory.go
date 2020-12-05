package GoogleAdminDirectory

import (
	bigquerytools "github.com/leapforce-libraries/go_bigquerytools"
	google "github.com/leapforce-libraries/go_google"
)

const (
	apiName string = "GoogleAdminDirectory"
	apiURL  string = "https://www.googleapis.com/admin/directory/v1"
)

// GoogleAdminDirectory stores GoogleAdminDirectory configuration
//
type GoogleAdminDirectory struct {
	Client *google.GoogleClient
}

// methods
//
func NewGoogleAdminDirectory(clientID string, clientSecret string, scope string, bigQuery *bigquerytools.BigQuery) *GoogleAdminDirectory {
	config := google.GoogleClientConfig{
		APIName:      apiName,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scope:        scope,
	}

	googleClient := google.NewGoogleClient(config, bigQuery)

	return &GoogleAdminDirectory{googleClient}
}
