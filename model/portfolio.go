package model

type Portfolio struct {
	ID             int    `json:"id" form:"id"`
	Title          string `json:"title" form:"title"`
	Description    string `json:"description" form:"description"`
	Category       string `json:"category" form:"category"`
	SEOTags        string `json:"seo_tags" form:"seo_tags"`
	SEODescription string `json:"seo_description" form:"seo_description"`
	Slug           string `json:"slug" form:"slug"`
}
