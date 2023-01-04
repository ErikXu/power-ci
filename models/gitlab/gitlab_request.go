package gitlab

type GitlabRequest interface {
	EmptyRequest | OauthRequest | CreateUserRequest |
		CreatePersonalAccessTokenRequest | CreateProjectRequest
}
