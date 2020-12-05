package GoogleAdminDirectory

import (
	"net/http"

	bigquerytools "github.com/leapforce-libraries/go_bigquerytools"
	google "github.com/leapforce-libraries/go_google"
)

const (
	apiName         string = "GoogleAdminDirectory"
	apiURL          string = "https://www.googleapis.com/admin/directory/v1"
	authURL         string = "https://accounts.google.com/o/oauth2/v2/auth"
	tokenURL        string = "https://oauth2.googleapis.com/token"
	tokenHTTPMethod string = http.MethodPost
	redirectURL     string = "http://localhost:8080/oauth/redirect"
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
