package endpoints

import (
	ReturnModule "ba-digital/backend/modules/return_module"
	"ba-digital/backend/structs"
	"net/http"
	"strconv"
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
	name := r.Form.Get("name")
	class := r.Form.Get("class")
	if name == "" || class == "" {
		ReturnModule.CustomError(w, r, "Bad Request: name and class cannot be empty", 400)
		return
	}
	score, error := strconv.Atoi(r.Form.Get("score"))
	if error != nil {
		ReturnModule.CustomError(w, r, "Bad Request: score must be an int", 400)
		return
	}
	newUser := structs.User{Name: name, Class: class, Score: score}
	structs.Users = append(structs.Users, newUser)
	ReturnModule.Success(w, r)
}
