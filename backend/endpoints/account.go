package endpoints

import (
	ReturnModule "ba-digital/backend/modules/return_module"
	"ba-digital/backend/structs"
	"net/http"
)

func SubmitResult(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		postRequest(w, r)
	} else {
		ReturnModule.MethodNotAllowed(w, r)
	}
}

func postRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	newUser := structs.UserResponse{Username: username, Ranking: 3}
	structs.Users = append(structs.Users, newUser)
}
