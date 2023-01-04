package gitlab

type GetNamespaceItem struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Kind      string `json:"kind"`
	FullPath  string `json:"full_path"`
	ParentId  int    `json:"parent_id"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

type GetNamespaceResponse = []GetNamespaceItem
