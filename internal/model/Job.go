package model

import (
	"time"
)

const (
	Remote = iota
	Hybrid
	Office
)

type Job struct {
	ID           int64    `json:"id"`
	RoleName     string   `json:"roleName" validate:"required,alpha"`
	Description  string   `json:"description" validate:"required"`
	Experience   int      `json:"experience" validate:"required,gte=0"`
	Locations    []string `json:"locations" validate:"alpha"`
	RemoteStatus int      `json:"remoteStatus" validate:"required,number,gt=0,lte=2"`
	SalaryRange  []int    `json:"salaryRange" validate:"number,gt=0"`
	CreatedOn    string   `json:"-"`
	UpdatedOn    string   `json:"-"`
	DeleteOn     string   `json:"-"`
}

type Jobs []*Job

func (j *Job) Validate() {
}

var JobList = []*Job{
	&Job{
		ID:           1,
		RoleName:     "Software Engineer",
		Description:  "Should be experienced with Python, Django, and SQL tech stack",
		Experience:   3,
		RemoteStatus: Hybrid,
		SalaryRange:  []int{30000, 50000},
		CreatedOn:    time.Now().UTC().String(),
	},
	&Job{
		ID:           2,
		RoleName:     "Senior Frontend Engineer",
		Description:  "Should be cracked dev in TypeScript, Svelte, Vercel stack",
		Experience:   3,
		RemoteStatus: Remote,
		SalaryRange:  []int{100000, 200000},
		CreatedOn:    time.Now().UTC().String(),
	},
}

func AddJob(req Job) {
	JobList = append(JobList, &req)
}

func DeleteJob(id int) {
	JobList = append(JobList[:id], JobList[id+1:]...)
}
