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
	score, _ := strconv.Atoi(r.Form.Get("score")) //TODO Handle error
	newUser := structs.User{Name: name, Class: class, Score: score}
	structs.Users = append(structs.Users, newUser)
}
