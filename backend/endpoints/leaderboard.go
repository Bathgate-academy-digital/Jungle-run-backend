package endpoints

import (
	"ba-digital/backend/database"
	ReturnModule "ba-digital/backend/modules/return_module"
	"net/http"
)

func GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getRequest(w, r)
	} else {
		ReturnModule.MethodNotAllowed(w, r)
	}
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	ReturnModule.Leaderboard(w, r, database.GetLeaderboard())
}
