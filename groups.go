package GoogleAdminDirectory

import (
	"fmt"
	"net/url"
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

func (gad *GoogleAdminDirectory) Groups(domain string) (*[]Group, error) {

	err := gad.Validate()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%sgroups?domain=%s", gad.baseURL, url.QueryEscape(domain))
	//fmt.Println(url)

	groupsReponse := GroupsResponse{}

	_, err = gad.oAuth2.Get(url, &groupsReponse)
	if err != nil {
		return nil, err
	}

	return &groupsReponse.Groups, nil
}
