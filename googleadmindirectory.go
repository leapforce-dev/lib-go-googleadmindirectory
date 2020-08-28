package GoogleAdminDirectory

import (
	"fmt"
	"strings"

	bigquerytools "github.com/Leapforce-nl/go_bigquerytools"
	types "github.com/Leapforce-nl/go_types"

	googleoauth2 "github.com/Leapforce-nl/go_googleoauth2"
)

const apiName string = "GoogleAdminDirectory"

// GoogleAdminDirectory stores GoogleAdminDirectory configuration
//
type GoogleAdminDirectory struct {
	baseURL string
	oAuth2  *googleoauth2.GoogleOAuth2
}

// methods
//
func NewGoogleAdminDirectory(baseURL string, clientID string, clientSecret string, scopes []string, bigQuery *bigquerytools.BigQuery, isLive bool) (*GoogleAdminDirectory, error) {
	gsc := GoogleAdminDirectory{}
	gsc.baseURL = baseURL

	_oAuth2 := new(googleoauth2.GoogleOAuth2)
	_oAuth2.ApiName = apiName
	_oAuth2.ClientID = clientID
	_oAuth2.ClientSecret = clientSecret
	_oAuth2.Scopes = scopes
	_oAuth2.BigQuery = bigQuery
	_oAuth2.IsLive = isLive

	gsc.oAuth2 = _oAuth2

	return &gsc, nil
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
