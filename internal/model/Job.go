package model

import "time"

type Job struct {
	ID          int64  `json:"id"`
	RoleName    string `json:"roleName"`
	Description string `json:"description"`
	Experience  int    `json:"experience"`
	CreatedOn   string `json:"-"`
	UpdatedOn   string `json:"-"`
	DeleteOn    string `json:"-"`
}

type Jobs []*Job

var JobList = []*Job{
	&Job{
		ID:          1,
		RoleName:    "Software Engineer",
		Description: "Should be experienced with Python, Django, and SQL tech stack",
		Experience:  3,
		CreatedOn:   time.Now().UTC().String(),
	},
	&Job{
		ID:          2,
		RoleName:    "Senior Frontend Engineer",
		Description: "Should be cracked dev in TypeScript, Svelte, Vercel stack",
		Experience:  3,
		CreatedOn:   time.Now().UTC().String(),
	},
}
