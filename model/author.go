package model

type Author struct {
	ID        int    `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Instagram string `json:"instagram" form:"instagram"`
	Facebook  string `json:"facebook" form:"facebook"`
	Linkedin  string `json:"linkedin" form:"linkedin"`
	Github    string `json:"github" form:"github"`
}
