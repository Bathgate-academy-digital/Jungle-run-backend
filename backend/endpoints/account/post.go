package account

import (
	"ba-digital/backend/structs"
	"net/http"
)

func PostRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	// ranking := r.Header.Get("ranking")
	newUser := structs.UserResponse{Username: username, Ranking: 3}
	structs.Users = append(structs.Users, newUser)
}
