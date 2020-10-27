package GoogleAdminDirectory

import (
	"fmt"
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

func (gad *GoogleAdminDirectory) Members(groupID string) (*[]Member, error) {

	err := gad.Validate()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%sgroups/%s/members", gad.baseURL, groupID)
	//fmt.Println(url)

	membersReponse := MembersResponse{}

	_, err = gad.oAuth2.Get(url, &membersReponse)
	if err != nil {
		return nil, err
	}

	return &membersReponse.Members, nil
}