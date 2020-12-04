package GoogleAdminDirectory

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	googleerror "github.com/leapforce-libraries/go_googleerror"
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

	googleError := googleerror.GoogleError{}
	membersReponse := MembersResponse{}

	_, _, e := gad.oAuth2.Get(url, &membersReponse, &googleError)
	if e != nil {
		if googleError.Error.Message != "" {
			e.SetMessage(googleError.Error.Message)
		}
		return nil, e
	}

	return &membersReponse.Members, nil
}
