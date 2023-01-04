package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"power-ci/models/gitlab"
	"reflect"
)

type GitlabClient struct {
	Client       http.Client
	Host         string
	AccessToken  string
	PrivateToken string
}

func (client *GitlabClient) GrantOauthToken(grantType string, username string, password string) gitlab.OauthResponse {
	request := &gitlab.OauthRequest{
		GrantType: grantType,
		Username:  username,
		Password:  password,
	}

	return call[gitlab.OauthRequest, gitlab.OauthResponse](client, "POST", "/oauth/token", request)
}

func (client *GitlabClient) GetUserByUsername(username string) gitlab.GetUserResponse {
	request := &gitlab.EmptyRequest{}
	return call[gitlab.EmptyRequest, gitlab.GetUserResponse](client, "GET", fmt.Sprintf("/api/v4/users?username=%s", username), request)
}

func (client *GitlabClient) GetNamespaces() gitlab.GetNamespaceResponse {
	request := &gitlab.EmptyRequest{}
	return call[gitlab.EmptyRequest, gitlab.GetNamespaceResponse](client, "GET", "/api/v4/namespaces", request)
}

func (client *GitlabClient) CreateUser(admin bool, username string, name string, email string, password string, skipConfirmation bool) gitlab.CreateUserResponse {
	request := &gitlab.CreateUserRequest{
		Admin:            admin,
		Username:         username,
		Name:             name,
		Email:            email,
		Password:         password,
		SkipConfirmation: skipConfirmation,
	}

	return call[gitlab.CreateUserRequest, gitlab.CreateUserResponse](client, "POST", "/api/v4/users", request)
}

func (client *GitlabClient) CreatePersonalAccessToken(userId int, name string, scopes []string, expiresAt string) gitlab.CreatePersonalAccessTokenResponse {
	request := &gitlab.CreatePersonalAccessTokenRequest{
		Name:      name,
		Scopes:    scopes,
		ExpiresAt: expiresAt,
	}

	return call[gitlab.CreatePersonalAccessTokenRequest, gitlab.CreatePersonalAccessTokenResponse](client, "POST", fmt.Sprintf("/api/v4/users/%d/personal_access_tokens", userId), request)
}

func (client *GitlabClient) CreateProject(name string, namespaceId int) gitlab.CreateProjectResponse {
	request := &gitlab.CreateProjectRequest{
		Name:        name,
		NamespaceId: namespaceId,
	}

	return call[gitlab.CreateProjectRequest, gitlab.CreateProjectResponse](client, "POST", "/api/v4/projects", request)
}

func callToBytes[Request gitlab.GitlabRequest](client *GitlabClient, method string, api string, request *Request) []byte {
	url := client.Host + api

	req, _ := http.NewRequest(method, url, nil)

	if reflect.TypeOf(request) != reflect.TypeOf(gitlab.EmptyRequest{}) {
		jsonValue, _ := json.Marshal(request)
		req, _ = http.NewRequest(method, url, bytes.NewBuffer(jsonValue))
	}

	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	if client.AccessToken != "" {
		req.Header.Add("Authorization", "Bearer "+client.AccessToken)
	}

	if client.PrivateToken != "" {
		req.Header.Add("PRIVATE-TOKEN", client.PrivateToken)
	}

	res, err := client.Client.Do(req)
	if err != nil {
		panic("Call rest api failed")
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	return body
}

func call[Request gitlab.GitlabRequest, Response gitlab.GitlabResponse](client *GitlabClient, method string, api string, request *Request) Response {
	body := callToBytes(client, method, api, request)
	response := new(Response)
	json.Unmarshal(body, &response)
	return *response
}
