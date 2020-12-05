package GoogleAdminDirectory

import (
	"fmt"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
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

func (gad *GoogleAdminDirectory) Groups(domain string) (*[]Group, *errortools.Error) {
	url := fmt.Sprintf("%s/groups?domain=%s", apiURL, url.QueryEscape(domain))

	groupsReponse := GroupsResponse{}

	_, _, e := gad.Client.Get(url, &groupsReponse)
	if e != nil {
		return nil, e
	}

	return &groupsReponse.Groups, nil
}
