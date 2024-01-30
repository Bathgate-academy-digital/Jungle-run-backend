package endpoints

import (
	"jungle-rush/backend/database"
	ReturnModule "jungle-rush/backend/modules/return_module"
	"jungle-rush/backend/structs"
	"net/http"
)

func GetAdminLeaderboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ReturnModule.MethodNotAllowed(w)
		return
	}

	class := r.FormValue("class") // Extract the class parameter from the request

	if class != "" && len(class) != 3 {
		ReturnModule.BadRequest(w, "Invalid class parameter")
		return
	}

	var result []structs.User
	if class == "" {
		// no class specified, fetch the overall leaderboard
		result = database.GetAdminLeaderboard()
	} else {
		// class specified, fetch the leaderboard for that specific class
		result = database.GetAdminLeaderboardByClass(class)
	}

	if result == nil {
		ReturnModule.InternalServerError(w, "Error fetching leaderboard")
		return
	}

	ReturnModule.Leaderboard(w, result)
}
