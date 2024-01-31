package endpoints

import (
	"jungle-rush/backend/database"
	AdminModule "jungle-rush/backend/modules/admin_module"
	ReturnModule "jungle-rush/backend/modules/return_module"
	"jungle-rush/backend/structs"
	"log"
	"net/http"
)

func GetAdminLeaderboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ReturnModule.MethodNotAllowed(w)
		return
	}
	if !AdminModule.IsCorrectCreds(r) {
		ReturnModule.Unauthorized(w, "Invalid credentials")
		return
	}

	username, password, ok := r.BasicAuth()
	if username != "admin" || password != "5fe88ee2442925b67c5aa328ae3c65445b66d195fef92e705360931d3c2f037b" || !ok {
		log.Printf("Wrong credentials: username=%v password=%v", username, password)
		ReturnModule.Unauthorized(w, "Invalid credentials")
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
