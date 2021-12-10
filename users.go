package googleadmindirectory

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type UsersResponse struct {
	Kind  string `json:"kind"`
	Etag  string `json:"etag"`
	Users []User `json:"users"`
}

type User struct {
	Kind         string `json:"kind"`
	Etag         string `json:"etag"`
	ID           string `json:"id"`
	PrimaryEmail string `json:"primaryEmail"`
	Name         struct {
		GivenName  string `json:"givenName"`
		FamilyName string `json:"familyName"`
		FullName   string `json:"fullName"`
	} `json:"name"`
	IsAdmin                   bool   `json:"isAdmin"`
	IsDelegatedAdmin          bool   `json:"isDelegatedAdmin"`
	Status                    string `json:"status"`
	LastLoginTime             string `json:"lastLoginTime"`
	CreationTime              string `json:"creationTime"`
	AgreedToTerms             bool   `json:"agreedToTerms"`
	Suspended                 bool   `json:"suspended"`
	Archived                  bool   `json:"archived"`
	ChangePasswordAtNextLogin bool   `json:"changePasswordAtNextLogin"`
	IPWhitelisted             bool   `json:"ipWhitelisted"`
	Emails                    []struct {
		Address string `json:"address"`
		Type    string `json:"type"`
		Primary bool   `json:"primary"`
	} `json:"emails"`
	Phones []struct {
		Value string `json:"value"`
		Type  string `json:"type"`
	} `json:"phones"`
	NonEditableAliases         []string `json:"nonEditableAliases"`
	CustomerID                 string   `json:"customerId"`
	OrgUnitPath                string   `json:"orgUnitPath"`
	IsMailboxSetup             bool     `json:"isMailboxSetup"`
	IsEnrolledIn2SV            bool     `json:"isEnrolledIn2Sv"`
	IsEnforcedIn2Sv            bool     `json:"isEnforcedIn2Sv"`
	IncludeInGlobalAddressList bool     `json:"includeInGlobalAddressList"`
	ThumbnailPhotoURL          string   `json:"thumbnailPhotoUrl"`
	ThumbnailPhotoEtag         string   `json:"thumbnailPhotoEtag"`
}

func (service *Service) Users(domain string) (*[]User, *errortools.Error) {
	usersReponse := UsersResponse{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		URL:           service.url(fmt.Sprintf("users?domain=%s", domain)),
		ResponseModel: &usersReponse,
	}

	_, _, e := service.googleService.HTTPRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &usersReponse.Users, nil
}
