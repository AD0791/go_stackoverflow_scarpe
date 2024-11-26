package models

type Question struct {
	Title   string `json:"title"`
	User    string `json:"user"`
	Votes   int    `json:"votes"`
	Answers int    `json:"answers"`
	Views   int    `json:"views"`
}

type QuestionsResponse struct {
	Questions []Question `json:"questions"`
	Total     int        `json:"total"`
}
