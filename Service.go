package googleadmindirectory

import (
	errortools "github.com/leapforce-libraries/go_errortools"
	google "github.com/leapforce-libraries/go_google"
)

const (
	apiName string = "GoogleAdminDirectory"
	apiURL  string = "https://www.googleapis.com/admin/directory/v1"
)

// GoogleAdminDirectory stores GoogleAdminDirectory configuration
//
type Service struct {
	googleService *google.Service
}

// methods
//
func NewService(clientID string, clientSecret string, scope string, bigQuery *google.BigQuery) *Service {
	config := google.ServiceConfig{
		APIName:      apiName,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scope:        scope,
	}

	googleService := google.NewService(config, bigQuery)

	return &Service{googleService}
}

func (service *Service) InitToken() *errortools.Error {
	return service.googleService.InitToken()
}
