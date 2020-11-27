package models

type User struct {
	ID              int64  `json:"ID"`
	Timestamp       string `json:"Timestamp"`
	ID2             string `json:"ID2"`
	UID             string `json:"UID"`
	Name            string `json:"Name"`
	Gender          string `json:"Gender"`
	Email           string `json:"Email"`
	Phone           string `json:"Phone"`
	Dept            string `json:"Dept"`
	Grade           string `json:"Grade"`
	Language        string `json:"Language"`
	Region          string `json:"Region"`
	Role            string `json:"Role"`
	PreferTags      string `json:"PreferTags"`
	ObtainedCredits string `json:"ObtainedCredits"`
}
