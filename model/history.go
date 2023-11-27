package model

type History struct {
	Id        int    `json:"id"`
	Url       string `json:"url"`
	Updated   string `json:"updated"`
	IsDeleted bool   `json:"isDeleted"`
	Note      string `json:"note"`
}

type QSHistory struct {
	Id      int    `json:"id"`
	Url     string `json:"url"`
	Qs_name string `json:"qs_name"`
}

type ContentText struct {
	Id      int    `json:"id"`
	Text    string `json:"text"`
	Updated string `json:"updated"`
}
