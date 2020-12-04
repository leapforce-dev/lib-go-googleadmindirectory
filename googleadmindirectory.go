package GoogleAdminDirectory

import (
	"net/http"

	bigquerytools "github.com/leapforce-libraries/go_bigquerytools"
	errortools "github.com/leapforce-libraries/go_errortools"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
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
	oAuth2 *oauth2.OAuth2
}

// methods
//
func NewGoogleAdminDirectory(clientID string, clientSecret string, scope string, bigQuery *bigquerytools.BigQuery) *GoogleAdminDirectory {
	gad := GoogleAdminDirectory{}

	config := oauth2.OAuth2Config{
		ApiName:         apiName,
		ClientID:        clientID,
		ClientSecret:    clientSecret,
		Scope:           scope,
		RedirectURL:     redirectURL,
		AuthURL:         authURL,
		TokenURL:        tokenURL,
		TokenHTTPMethod: tokenHTTPMethod,
	}
	gad.oAuth2 = oauth2.NewOAuth(config, bigQuery)

	return &gad
}

func (gad *GoogleAdminDirectory) InitToken() *errortools.Error {
	return gad.oAuth2.InitToken()
}
