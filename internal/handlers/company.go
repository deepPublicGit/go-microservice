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
		getJobs(rw)
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
