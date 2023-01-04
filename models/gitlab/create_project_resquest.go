package gitlab

type CreateProjectRequest struct {
	Name        string `json:"name"`
	NamespaceId int    `json:"namespace_id"`
}
