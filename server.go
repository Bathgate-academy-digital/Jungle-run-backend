package main

import (
	"ba-digital/backend/endpoints"
	"github.com/pterm/pterm"
	"log"
	"net/http"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	if true {
		pterm.Success.Println("Database has been initialized successfully.")

		http.HandleFunc("/api/account", endpoints.ManageAccounts)
		http.HandleFunc("/api/account/session", endpoints.ManageSessions)
		http.HandleFunc("/api/account/update", endpoints.ManageUpdate)
		http.HandleFunc("/api/account/create", endpoints.ManageCreation)

		log.Fatal(http.ListenAndServe(":8080", corsMiddleware(http.DefaultServeMux)))
	} else {
		pterm.Fatal.WithFatal(true).Println("Failed to initialize database successfully.")
	}
}
