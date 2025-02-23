package model

import (
	"time"
)

type Company struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name" validate:"required"`
	Slogan      string        `json:"slogan"`
	Description string        `json:"description"`
	LogoURL     string        `json:"logoURL" validate:"url"`
	Webpage     string        `json:"webpage" validate:"url"`
	Batch       string        `json:"batch" validate:"required,alphanum"`
	Tags        []string      `json:"tags" validate:"alpha"`
	SocialMedia []SocialMedia `json:"socialMedia"`
	CompanySize int           `json:"companySize" validate:"required,number,gt=0"`
	Jobs        []*Job        `json:"jobs"`
	CreatedOn   string        `json:"-"`
	UpdatedOn   string        `json:"-"`
	DeletedOn   string        `json:"-"`
}

type SocialMedia struct {
	ID         int64  `json:"id"`
	UserName   string `json:"userName"`
	Service    string `json:"service"`
	ProfileURL string `json:"profileURL"`
	CreatedOn  string `json:"-"`
	UpdatedOn  string `json:"-"`
	DeletedOn  string `json:"-"`
}

type Companies []*Company

func (c *Company) Validate() {
}

var CompanyList = []*Company{
	&Company{
		ID:          1,
		Name:        "Alpha",
		Slogan:      "Making alpha greater than beta",
		Description: "We are global disruptors who want to make the next great company and enable alpha to be greater than beta.",
		LogoURL:     "/static/images/alphaLogo.png",
		Webpage:     "https://localhost:9090/about",
		Batch:       "W2025",
		Tags:        []string{"Engineering", "Stocks & Bonds"},
		SocialMedia: getDummySocialMedia("alpha"),
		CompanySize: 10,
		Jobs:        JobList,
		CreatedOn:   time.Now().UTC().Format(time.RFC3339),
		UpdatedOn:   time.Now().UTC().Format(time.RFC3339),
		DeletedOn:   time.Now().UTC().Format(time.RFC3339),
	},
	&Company{
		ID:          2,
		Name:        "Beta",
		Slogan:      "Making beta greater than alpha",
		Description: "We are global disruptors who dont want to make the next great company and enable beta to be greater than alpha.",
		LogoURL:     "/static/images/betaLogo.png",
		Webpage:     "https://localhost:9091/about",
		Batch:       "S2024",
		Tags:        []string{"Retail", "Finance and Accounting"},
		SocialMedia: getDummySocialMedia("beta"),
		CompanySize: 10,
		CreatedOn:   time.Now().UTC().Format(time.RFC3339),
		UpdatedOn:   time.Now().UTC().Format(time.RFC3339),
		DeletedOn:   time.Now().UTC().Format(time.RFC3339),
	},
}

func getDummySocialMedia(name string) []SocialMedia {
	DummySocials := make([]SocialMedia, 0, 2)
	DummySocials = append(DummySocials, SocialMedia{
		ID:         1,
		UserName:   name,
		Service:    "instagram",
		ProfileURL: "https://www.instagram.com/" + name,
		CreatedOn:  time.Now().UTC().Format(time.RFC3339),
		UpdatedOn:  time.Now().UTC().Format(time.RFC3339),
		DeletedOn:  time.Now().UTC().Format(time.RFC3339)})
	DummySocials = append(DummySocials, SocialMedia{
		ID:         2,
		UserName:   name,
		Service:    "twitter",
		ProfileURL: "https://www.twitter.com/" + name,
		CreatedOn:  time.Now().UTC().Format(time.RFC3339),
		UpdatedOn:  time.Now().UTC().Format(time.RFC3339),
		DeletedOn:  time.Now().UTC().Format(time.RFC3339)})
	return DummySocials
}

func AddCompany(req Company) {
	CompanyList = append(CompanyList, &req)
}

func DeleteCompany(id int) {
	CompanyList = append(CompanyList[:id], CompanyList[id+1:]...)
}
