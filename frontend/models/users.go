package models

type User struct {
	ID              int64  `json:"id"`
	Timestamp       int64  `json:"timestamp"`
	UID             string `json:"uid"`
	Name            string `json:"name"`
	Gender          string `json:"gender"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Dept            string `json:"dept"`
	Grade           string `json:"grade"`
	Language        string `json:"language"`
	Region          string `json:"region"`
	Role            string `json:"role"`
	PreferTags      string `json:"prefer_tags"`
	ObtainedCredits string `json:"obtained_credits"`
}
