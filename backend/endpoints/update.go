package endpoints

import (
	"jungle-rush/backend/database"
	ReturnModule "jungle-rush/backend/modules/return_module"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ReturnModule.MethodNotAllowed(w)
		return
	}

	r.ParseForm()
	idStr := r.Form.Get("id")
	timeStr := r.Form.Get("time")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ReturnModule.BadRequest(w, "Invalid id")
		return
	}

	err = database.UpdateUser(id, timeStr)
	if err != nil {
		ReturnModule.InternalServerError(w, "Failed to update user")
		return
	}

	ReturnModule.Success(w, r)
}
