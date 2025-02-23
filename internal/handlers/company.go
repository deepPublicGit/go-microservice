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

func (s *Companies) GetCompaniesByID(rw http.ResponseWriter, req *http.Request) {
	println("GET RECEIVED")
	encoder := json.NewEncoder(rw)
	err := encoder.Encode(model.CompanyList)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

}

func (s *Companies) GetJobsByCompany(rw http.ResponseWriter, req *http.Request) {
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

func (s *Companies) DeleteCompanies(rw http.ResponseWriter, req *http.Request) {
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

	model.DeleteCompany(id)
	rw.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(rw, "Deleted %d Successfully", id)
}
