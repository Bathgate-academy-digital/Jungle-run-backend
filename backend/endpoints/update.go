package endpoints

import (
	"ba-digital/backend/database"
	ReturnModule "ba-digital/backend/modules/return_module"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ReturnModule.MethodNotAllowed(w, r)
		return
	}

	r.ParseForm()
	idStr := r.Form.Get("id")
	scoreStr := r.Form.Get("score")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ReturnModule.CustomError(w, r, "Invalid id", 400)
		return
	}
	score, err := strconv.Atoi(scoreStr)
	if err != nil {
		ReturnModule.CustomError(w, r, "Invalid score", 400)
		return
	}

	err = database.UpdateUser(id, score)
	if err != nil {
		ReturnModule.CustomError(w, r, "Failed to update user", 500)
		return
	}

	ReturnModule.Success(w, r)
}
