package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/deepPublicGit/go-microservice/internal/model"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"strconv"
)

type Jobs struct {
	l *slog.Logger
}

func NewJobs(l *slog.Logger) *Jobs {
	return &Jobs{l: l}
}

func (s *Jobs) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.l.Info("YOLO RECEIVED", req.Header)
	if req.Method == "GET" {
		s.GetJobs(rw, req)
	}
}

func (s *Jobs) GetJobsByID(rw http.ResponseWriter, req *http.Request) {
	println("GET RECEIVED")
	encoder := json.NewEncoder(rw)
	err := encoder.Encode(model.JobList)
	if err != nil {
		return
	}
}

func (s *Jobs) GetJobsByCompany(rw http.ResponseWriter, req *http.Request) {
	println("GET RECEIVED")
	encoder := json.NewEncoder(rw)
	err := encoder.Encode(model.JobList)
	if err != nil {
		return
	}
}

func (s *Jobs) GetJobs(rw http.ResponseWriter, req *http.Request) {
	println("GET RECEIVED")
	encoder := json.NewEncoder(rw)
	err := encoder.Encode(model.JobList)
	if err != nil {
		return
	}
}

func (s *Jobs) AddJobs(rw http.ResponseWriter, req *http.Request) {
	println("POST RECEIVED")
	decoder := json.NewDecoder(req.Body)
	_, err := decoder.Token()
	if err != nil {
		s.l.Info("ERROR", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	if decoder.More() {
		var r model.Job
		err := decoder.Decode(&r)
		if err != nil {
			s.l.Info("ERROR2", err)

			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		s.l.Info("AddJobs", r)
		model.AddJob(r)
	}
	_, err = decoder.Token()
	if err != nil {
		s.l.Error("YOLO ERROR", err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Jobs) DeleteJobs(rw http.ResponseWriter, req *http.Request) {
	println("DELETE RECEIVED")
	vars := mux.Vars(req)

	if vars["id"] == "" {
		http.Error(rw, "Empty ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}

	model.DeleteJob(id - 1)
	rw.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(rw, "Deleted %d Successfully", id)
}
