package model

import (
	"testing"
)

func TestCompany_ValidateFail(t *testing.T) {
	c := &Company{
		Name: "",
	}
	err := c.Validate()
	if err == nil {
		t.Fatal("want validation error", err)
	}
}

func TestCompany_ValidatePass(t *testing.T) {
	c := &Company{
		Name:        "Sample",
		Batch:       "W2025",
		CompanySize: 25,
	}
	err := c.Validate()
	if err != nil {
		t.Fatal("Want no validation failure", err)
	}
}
