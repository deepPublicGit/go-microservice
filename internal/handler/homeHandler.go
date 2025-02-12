package handler

import (
	"fmt"
	"io"
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
	b, _ := io.ReadAll(req.Body)

	fmt.Fprintf(rw, "YOLO %s", b)
}
