package model

import "time"

type Job struct {
	ID          int64
	RoleName    string
	Description string
	LogoURL     string
	Experience  int
	CreatedOn   string
	DeleteOn    string
}

var JobList = []*Job{
	&Job{
		ID:          1,
		RoleName:    "Software Engineer",
		Description: "Should be experienced with Python, Django, and SQL tech stack",
		LogoURL:     "https://localhost:8080/getDemoIcon.ico",
		Experience:  3,
		CreatedOn:   time.Now().UTC().String(),
	},
	&Job{
		ID:          2,
		RoleName:    "Senior Frontend Engineer",
		Description: "Should be cracked dev in TypeScript, Svelte, Vercel stack",
		LogoURL:     "https://localhost:8080/getDemoIcon.ico",
		Experience:  3,
		CreatedOn:   time.Now().UTC().String(),
	},
}
