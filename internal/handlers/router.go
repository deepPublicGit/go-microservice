package handlers

import (
	"github.com/gorilla/mux"
	"log/slog"
)

// Get, Get Batch, Post, Post Batch, Later Get Pagination, Patch single.
// Post Company, Get Company,Companies,Put,Patch

func RouterHandler(logger *slog.Logger) *mux.Router {
	router := mux.NewRouter()
	router.Handle("/", NewHomeHandler(logger))
	jh := NewJobs(logger)

	jobsGetRouter := router.PathPrefix("/jobs").Methods("GET").Subrouter()
	jobsGetRouter.HandleFunc("/", jh.GetJobs)
	jobsGetRouter.HandleFunc("/{id}", jh.GetJobsByID)
	jobsGetRouter.HandleFunc("/{id:[0-9]+}/company", jh.GetJobsByCompany)

	jobsPostRouter := router.PathPrefix("/jobs").Methods("POST").Subrouter()
	jobsPostRouter.HandleFunc("/", jh.AddJobs)

	jobsDeleteRouter := router.PathPrefix("/jobs").Methods("DELETE").Subrouter()
	jobsDeleteRouter.HandleFunc("/{id:[0-9]+}", jh.DeleteJobs)

	ch := NewCompanies(logger)
	companiesRouter := router.PathPrefix("/companies").Methods("GET").Subrouter()
	companiesRouter.HandleFunc("/", ch.GetCompanies)
	companiesRouter.HandleFunc("/{id}", ch.GetCompaniesByID)
	companiesRouter.HandleFunc("/{company}/jobs", ch.GetJobsByCompany)

	companiesPostRouter := router.PathPrefix("/companies").Methods("POST").Subrouter()
	companiesPostRouter.HandleFunc("/", ch.AddCompanies)

	companiesDeleteRouter := router.PathPrefix("/jobs").Methods("DELETE").Subrouter()
	companiesDeleteRouter.HandleFunc("/{id:[0-9]+}", ch.DeleteCompanies)

	return router
}

/*	homeHandler := handlers.NewJobs(logger)
	servMux := http.NewServeMux()
	servMux.Handle("/", homeHandler)

	server := &http.Server{
		Addr:    ":8080",
		Companies: servMux,
	}
	err := server.ListenAndServe()
	if err != nil {
		logger.ErrorContext(context.Background(), "an error occurred", slog.String("error", err.Error()))
	}*/
