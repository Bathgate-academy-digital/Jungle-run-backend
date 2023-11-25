package endpoints

import (
	ReturnModule "ba-digital/backend/modules/return_module"
	"ba-digital/backend/structs"
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
	ReturnModule.Leaderboard(w, r, structs.Users)
}
