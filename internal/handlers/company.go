package handlers

import (
	"encoding/json"
	"github.com/deepPublicGit/go-microservice/internal/model"
	"log/slog"
	"net/http"
)

type Companies struct {
	l *slog.Logger
}

func NewCompanies(l *slog.Logger) *Companies {
	return &Companies{l: l}
}

func (s *Companies) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.l.Info("YOLO RECEIVED", req.Header)
	if req.Method == "GET" {
		s.GetCompanies(rw, req)
	}
}

func (s *Companies) GetCompanies(rw http.ResponseWriter, req *http.Request) {
	println("GET RECEIVED")
	encoder := json.NewEncoder(rw)
	err := encoder.Encode(model.CompanyList)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

}

func (s *Companies) AddCompanies(rw http.ResponseWriter, req *http.Request) {
	println("POST RECEIVED")
	decoder := json.NewDecoder(req.Body)
	_, err := decoder.Token()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	if decoder.More() {
		var r model.Company
		err := decoder.Decode(&r)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		s.l.Info("AddCompanies", r)
		model.AddCompany(r)
		s.l.Info("AddCompanies", model.CompanyList)
	}
	_, err = decoder.Token()
	if err != nil {
		s.l.Error("YOLO ERROR", err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
