package model

type Response struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Records []Records `json:"records"`
}

type Records struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	Info  string `json:"info"`
}
