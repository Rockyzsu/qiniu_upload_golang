package model

type History struct {
	Id        int    `json:"Id"`
	Url       string `json:"Url"`
	Updated   string `json:"Updated"`
	IsDeleted bool   `json:"isDeleted"`
}

type ContentText struct {
	Id      int    `json:"id"`
	Text    string `json:"text"`
	Updated string `json:"updated"`
}
