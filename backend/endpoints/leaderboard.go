package endpoints

import (
	"jungle-rush/backend/database"
	ReturnModule "jungle-rush/backend/modules/return_module"
	"net/http"
)

func GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ReturnModule.MethodNotAllowed(w, r)
		return
	}

	result := database.GetLeaderboard()
	if result == nil {
		ReturnModule.InternalServerError(w, r, "Error fetching leaderboard")
		return
	}
	ReturnModule.Leaderboard(w, r, result)
}
