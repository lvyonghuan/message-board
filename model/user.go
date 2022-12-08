package model

type User struct {
	UserName        string `json:"userName"`
	Password        string `json:"password"`
	SecrecyQuestion string `json:"secrecyQuestion"`
	Secrecy         string `json:"question"`
	Administrator   int    `json:"administrator"`
	Cookie          string `json:"cookie"`
}
