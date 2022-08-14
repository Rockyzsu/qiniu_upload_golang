package model

type History struct {
	Id        int    `json:"id"`
	Url       string `json:"url"`
	Updated   string `json:"updated"`
	IsDeleted bool   `json:"isDeleted"`
}

type ContentText struct {
	Id      int    `json:"id"`
	Text    string `json:"text"`
	Updated string `json:"updated"`
}
