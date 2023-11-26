package endpoints

import (
	"ba-digital/backend/database"
	ReturnModule "ba-digital/backend/modules/return_module"
	"math/rand"
	"net/http"
	"strconv"
)

func SubmitResult(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
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
		database.SubmitResult(rand.Int(), name, class, score)
		ReturnModule.Success(w, r)
	} else {
		ReturnModule.MethodNotAllowed(w, r)
	}

}
