package endpoints

import (
	"jungle-rush/backend/database"
	ReturnModule "jungle-rush/backend/modules/return_module"
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
		ReturnModule.BadRequest(w, r, "Invalid id")
		return
	}
	score, err := strconv.Atoi(scoreStr)
	if err != nil {
		ReturnModule.BadRequest(w, r, "Invalid score")
		return
	}

	err = database.UpdateUser(id, score)
	if err != nil {
		ReturnModule.InternalServerError(w, r, "Failed to update user")
		return
	}

	ReturnModule.Success(w, r)
}
