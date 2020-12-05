package GoogleAdminDirectory

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
)

type MembersResponse struct {
	Kind    string   `json:"kind"`
	Etag    string   `json:"etag"`
	Members []Member `json:"members"`
}

type Member struct {
	Kind   string `json:"kind"`
	Etag   string `json:"etag"`
	ID     string `json:"id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Type   string `json:"type"`
	Status string `json:"status"`
}

func (gad *GoogleAdminDirectory) Members(groupID string) (*[]Member, *errortools.Error) {
	url := fmt.Sprintf("%s/groups/%s/members", apiURL, groupID)

	membersReponse := MembersResponse{}

	_, _, e := gad.Client.Get(url, &membersReponse)
	if e != nil {
		return nil, e
	}

	return &membersReponse.Members, nil
}
