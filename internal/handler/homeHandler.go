package handler

import (
	"encoding/json"
	"github.com/deepPublicGit/go-microservice/internal/model"
	"log/slog"
	"net/http"
)

type Handler struct {
	l *slog.Logger
}

func NewHandler(l *slog.Logger) *Handler {
	return &Handler{l: l}
}

func (s *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.l.Info("YOLO RECEIVED", req.Header)
	if req.Method == "GET" {
		getJobs(rw)

	}
}

func getJobs(rw http.ResponseWriter) {
	println("GET RECEIVED")
	encoder := json.NewEncoder(rw)
	err := encoder.Encode(model.JobList)
	if err != nil {
		return
	}

}
