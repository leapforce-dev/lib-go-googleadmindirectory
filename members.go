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
	//fmt.Println(url)

	membersReponse := MembersResponse{}

	_, _, err := gad.oAuth2.Get(url, &membersReponse, nil)
	if err != nil {
		return nil, err
	}

	return &membersReponse.Members, nil
}
