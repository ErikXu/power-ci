package gitlab

import "time"

type CreateProjectResponse struct {
	Id                int       `json:"id"`
	Description       string    `json:"description"`
	Name              string    `json:"name"`
	NameWithNamespace string    `json:"name_with_namespace"`
	Path              string    `json:"path"`
	PathWithNamespace string    `json:"path_with_namespace"`
	CreatedAt         time.Time `json:"created_at"`
	DefaultBranch     string    `json:"default_branch"`
	SshUrlToRepo      string    `json:"ssh_url_to_repo"`
	HttpUrlToRepo     string    `json:"http_url_to_repo"`
	WebURL            string    `json:"web_url"`
}
