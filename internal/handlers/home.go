package handlers

import (
	"log/slog"
	"net/http"
)

type HomeHandler struct {
	l *slog.Logger
}

func NewHomeHandler(l *slog.Logger) *HomeHandler {
	return &HomeHandler{l: l}
}

func (s *HomeHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.l.Info("YOLO RECEIVED", req.Header)
	if req.Method == "GET" {
		getJobs(rw)

	}
}
