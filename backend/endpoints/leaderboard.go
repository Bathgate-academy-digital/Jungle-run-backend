package endpoints

import (
	"jungle-rush/backend/database"
	ReturnModule "jungle-rush/backend/modules/return_module"
	"net/http"
)

func GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ReturnModule.MethodNotAllowed(w)
		return
	}

	result := database.GetLeaderboard()
	if result == nil {
		ReturnModule.InternalServerError(w, "Error fetching leaderboard")
		return
	}
	ReturnModule.Leaderboard(w, result)
}
