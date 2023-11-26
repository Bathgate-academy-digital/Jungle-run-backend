package endpoints

import (
	"ba-digital/backend/database"
	ReturnModule "ba-digital/backend/modules/return_module"
	"math/rand"
	"net/http"
	"strconv"
)

func SubmitResult(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ReturnModule.MethodNotAllowed(w, r)
		return
	}

	r.ParseForm()
	name := r.Form.Get("name")
	class := r.Form.Get("class")

	if name == "" || class == "" {
		ReturnModule.CustomError(w, r, "Bad Request: name and class cannot be empty", 400)
		return
	}

	score, err := strconv.Atoi(r.Form.Get("score"))
	if err != nil {
		ReturnModule.CustomError(w, r, "Bad Request: score must be an int", 400)
		return
	}
	database.SubmitResult(rand.Int(), name, class, score)
	ReturnModule.Success(w, r)
}
