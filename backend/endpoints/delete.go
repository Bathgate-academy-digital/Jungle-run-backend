package endpoints

import (
	"fmt"
	"jungle-rush/backend/database"
	AdminModule "jungle-rush/backend/modules/admin_module"
	ReturnModule "jungle-rush/backend/modules/return_module"
	"net/http"
	"strconv"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ReturnModule.MethodNotAllowed(w)
		return
	}
	if !AdminModule.IsCorrectCreds(r) {
		ReturnModule.Unauthorized(w, "Invalid credentials")
		return
	}

	r.ParseForm()
	fmt.Println(r.Form)
	id := r.Form.Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ReturnModule.BadRequest(w, "id must be a number")
		return
	}

	result := database.DeleteUser(idInt)
	if result != nil {
		ReturnModule.InternalServerError(w, "Error deleting member from leaderboard")
		return
	}
	ReturnModule.Success(w, r)
}
