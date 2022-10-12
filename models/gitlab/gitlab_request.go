package gitlab

type GitlabRequest interface {
	EmptyRequest | OauthRequest
}
