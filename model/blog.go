package model

type Blog struct {
	ID             int    `json:"id" form:"id"`
	Title          string `json:"title" form:"title"`
	Description    string `json:"description" form:"description"`
	SEOTags        string `json:"seo_tags" form:"seo_tags"`
	SEODescription string `json:"seo_description" form:"seo_description"`
	ThumbnailsALT  string `json:"thumbnails_alt" form:"thumbnails_alt"`
	Thumbnails     string `json:"thumbnails" form:"thumbnails"`
	AuthorID       string `json:"author_id" form:"author_id"`
	CategoryID     string `json:"category_id" form:"category_id"`
	CreatedAt      string `json:"created_at" form:"created_at"`
}
