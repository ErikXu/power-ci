package gitlab

import "time"

type CreateCommitResponse struct {
	Id             string           `json:"id"`
	ShortId        string           `json:"short_id"`
	CreatedAt      time.Time        `json:"created_at"`
	Title          string           `json:"title"`
	Message        string           `json:"message"`
	AuthorName     string           `json:"author_name"`
	AuthorEmail    string           `json:"author_email"`
	AuthoredDate   time.Time        `json:"authored_date"`
	CommitterName  string           `json:"committer_name"`
	CommitterEmail string           `json:"committer_email"`
	CommittedDate  time.Time        `json:"committed_date"`
	WebURL         string           `json:"web_url"`
	Stats          CreateCommitStat `json:"stats"`
	ProjectId      int              `json:"project_id"`
}

type CreateCommitStat struct {
	Additions int `json:"additions"`
	Deletions int `json:"deletions"`
	Total     int `json:"total"`
}
