package model

type Company struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	Slogan      string        `json:"slogan"`
	Description string        `json:"description"`
	LogoURL     string        `json:"logoURL"`
	Webpage     string        `json:"webpage"`
	Batch       int           `json:"batch"`
	Tags        []string      `json:"tags"`
	SocialMedia []SocialMedia `json:"socialMedia"`
	CompanySize int           `json:"companySize"`
	CreatedOn   string        `json:"-"`
	UpdatedOn   string        `json:"-"`
	DeleteOn    string        `json:"-"`
}

type SocialMedia struct {
	ID         int64  `json:"id"`
	UserName   string `json:"userName"`
	Service    string `json:"service"`
	ProfileURL string `json:"profileURL"`
	CreatedOn  string `json:"-"`
	UpdatedOn  string `json:"-"`
	DeleteOn   string `json:"-"`
}

type Companies []*Company
