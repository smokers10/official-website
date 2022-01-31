package model

type Client struct {
	ID      int    `json:"id" form:"id"`
	Logo    string `json:"logo" form:"logo"`
	LogoALT string `json:"logo_alt" form:"logo_alt"`
	Name    string `json:"name" form:"name"`
}
