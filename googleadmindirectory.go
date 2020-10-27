package GoogleAdminDirectory

import (
	"fmt"
	"net/http"
	"strings"

	bigquerytools "github.com/Leapforce-nl/go_bigquerytools"
	types "github.com/Leapforce-nl/go_types"

	oauth2 "github.com/Leapforce-nl/go_oauth2"
)

const (
	apiName         string = "GoogleAdminDirectory"
	apiURL          string = "https://www.googleapis.com/calendar/v3"
	authURL         string = "https://accounts.google.com/o/oauth2/v2/auth"
	tokenURL        string = "https://oauth2.googleapis.com/token"
	tokenHTTPMethod string = http.MethodPost
	redirectURL     string = "http://localhost:8080/oauth/redirect"
)

// GoogleAdminDirectory stores GoogleAdminDirectory configuration
//
type GoogleAdminDirectory struct {
	baseURL string
	oAuth2  *oauth2.OAuth2
}

// methods
//
func NewGoogleAdminDirectory(clientID string, clientSecret string, scope string, bigQuery *bigquerytools.BigQuery, isLive bool) (*GoogleAdminDirectory, error) {
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
	gad.oAuth2 = oauth2.NewOAuth(config, bigQuery, isLive)

	return &gad, nil
}

func (gad *GoogleAdminDirectory) InitToken() error {
	return gad.oAuth2.InitToken()
}

func (gad *GoogleAdminDirectory) Validate() error {
	if gad.baseURL == "" {
		return &types.ErrorString{fmt.Sprintf("%s baseURL not provided", apiName)}
	}

	if !strings.HasSuffix(gad.baseURL, "/") {
		gad.baseURL = gad.baseURL + "/"
	}

	return nil
}
