package model

type Records struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	Info  string `json:"info"`
}

type RequestForm struct {
	Keyword string `form:"keyword"`
}
