package googleadmindirectory

import (
	"fmt"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

type GroupsResponse struct {
	Kind   string  `json:"kind"`
	Etag   string  `json:"etag"`
	Groups []Group `json:"groups"`
}

type Group struct {
	Kind               string   `json:"kind"`
	ID                 string   `json:"id"`
	Etag               string   `json:"etag"`
	Email              string   `json:"email"`
	Name               string   `json:"name"`
	DirectMembersCount string   `json:"directMembersCount"`
	Description        string   `json:"description"`
	AdminCreated       bool     `json:"adminCreated"`
	NonEditableAliases []string `json:"nonEditableAliases"`
}

func (service *Service) Groups(domain string) (*[]Group, *errortools.Error) {
	groupsResponse := GroupsResponse{}

	requestConfig := oauth2.RequestConfig{
		URL:           service.url(fmt.Sprintf("groups?domain=%s", url.QueryEscape(domain))),
		ResponseModel: &groupsResponse,
	}

	_, _, e := service.googleService.Get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &groupsResponse.Groups, nil
}
