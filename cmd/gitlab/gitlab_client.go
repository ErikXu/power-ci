package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"power-ci/models/gitlab"
)

type GitlabClient struct {
	Client      http.Client
	Host        string
	AccessToken string
}

func (client *GitlabClient) GrantOauthToken(grantType string, username string, password string) *gitlab.OauthResponse {
	request := &gitlab.OauthRequest{
		GrantType: grantType,
		Username:  username,
		Password:  password,
	}

	return call[gitlab.OauthRequest, gitlab.OauthResponse](client, "POST", "/oauth/token", *request)
}

func (client *GitlabClient) GetUserByUsername(username string) []gitlab.GetUserResponse {
	url := fmt.Sprintf("%s/api/v4/users?username=%s", client.Host, username)
	res, err := client.Client.Get(url)
	if err != nil {
		panic("Call rest api failed")
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	response := []gitlab.GetUserResponse{}
	json.Unmarshal(body, &response)
	return response
}

func (client *GitlabClient) CreateUser(admin bool, username string, name string, email string, password string) *gitlab.CreateUserResponse {
	request := &gitlab.CreateUserRequest{
		Admin:    admin,
		Username: username,
		Name:     name,
		Email:    email,
		Password: password,
	}

	return call[gitlab.CreateUserRequest, gitlab.CreateUserResponse](client, "POST", "/api/v4/users", *request)
}

func (client *GitlabClient) CreatePersonalAccessToken(userId int, name string, scopes []string, expiresAt string) *gitlab.CreatePersonalAccessTokenResponse {
	request := &gitlab.CreatePersonalAccessTokenRequest{
		Name:      name,
		Scopes:    scopes,
		ExpiresAt: expiresAt,
	}

	return call[gitlab.CreatePersonalAccessTokenRequest, gitlab.CreatePersonalAccessTokenResponse](client, "POST", fmt.Sprintf("/api/v4/users/%d/personal_access_tokens", userId), *request)
}

func callToBytes[Request gitlab.GitlabRequest](client *GitlabClient, method string, api string, request Request) []byte {
	url := client.Host + api

	jsonValue, _ := json.Marshal(request)
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(jsonValue))

	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	if client.AccessToken != "" {
		req.Header.Add("Authorization", "Bearer "+client.AccessToken)
	}

	res, err := client.Client.Do(req)
	if err != nil {
		panic("Call rest api failed")
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	return body
}

func call[Request gitlab.GitlabRequest, Response gitlab.GitlabResponse](client *GitlabClient, method string, api string, request Request) *Response {
	body := callToBytes(client, method, api, request)
	response := new(Response)
	json.Unmarshal(body, &response)
	return response
}
