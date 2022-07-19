package dto

type DeleteBook struct {
	SID string `json:"sid"`
}
type UpdateBook struct {
	SID         string `json:"sid"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type DBook struct {
	SID         string `json:"sid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserSID     string `json:"user_sid"`
	Path        string `json:"path"`
}
