package endpoints

import (
	"jungle-rush/backend/database"
	ReturnModule "jungle-rush/backend/modules/return_module"
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
	// Limit of PostgreSQL int https://www.postgresql.org/docs/15/datatype-numeric.html
	randomId := rand.Intn(2147483647)
	database.SubmitResult(randomId, name, class, score)
	ReturnModule.ID(w, r, randomId)
}
