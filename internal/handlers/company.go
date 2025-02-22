package handlers

import (
	"encoding/json"
	"github.com/deepPublicGit/go-microservice/internal/model"
	"log/slog"
	"net/http"
)

type CompanyHandler struct {
	l *slog.Logger
}

func NewCompanyHandler(l *slog.Logger) *CompanyHandler {
	return &CompanyHandler{l: l}
}

func (s *CompanyHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.l.Info("YOLO RECEIVED", req.Header)
	if req.Method == "GET" {
		getJobs(rw)

	}
}

func getCompanies(rw http.ResponseWriter) {
	println("GET RECEIVED")
	encoder := json.NewEncoder(rw)
	err := encoder.Encode(model.JobList)
	if err != nil {
		return
	}

}
