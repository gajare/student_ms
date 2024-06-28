package main

import (
	"log"
	"net/http"
	"student_ms/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Middleware
	r.Use(loggingMiddleware)

	// Routes
	
	r.HandleFunc("/students",handler.GetStudents).Methods("GET")
	r.HandleFunc("/students/{id}", handler.GetStudent).Methods("GET")
	r.HandleFunc("/students", handler.CreateStudent).Methods("POST")
	r.HandleFunc("/students/{id}", handler.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", handler.DeleteStudent).Methods("DELETE")

	// Start server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
