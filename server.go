package main

import (
	"jungle-rush/backend/database"
	"jungle-rush/backend/endpoints"
	"log"
	"net/http"

	"github.com/pterm/pterm"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://www.junglerush.co.uk")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	database.InitializeDatabase()
	database.CreateTables()

	pterm.Success.Println("Database has been initialized successfully.")

	http.HandleFunc("/api/create", endpoints.CreateUser)
	http.HandleFunc("/api/leaderboard", endpoints.GetLeaderboard)
	http.HandleFunc("/api/admin/leaderboard", endpoints.GetAdminLeaderboard)
	http.HandleFunc("/api/update", endpoints.Update)
	http.HandleFunc("/api/delete", endpoints.DeleteUser)

	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(http.DefaultServeMux)))
}
