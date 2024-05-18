package entity

type Score struct {
	UserID string `json:"userId"`
	Score  int    `json:"score"`
	Name   string `json:"name"`
}
