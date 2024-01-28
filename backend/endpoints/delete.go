package endpoints

import (
	"fmt"
	"jungle-rush/backend/database"
	ReturnModule "jungle-rush/backend/modules/return_module"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ReturnModule.MethodNotAllowed(w)
		return
	}

	r.ParseForm()
	fmt.Println(r.Form)
	name := r.Form.Get("name")
	class := r.Form.Get("class")
	if !validParams(w, r, name, class) {
		return
	}

	userId := database.GetUserId(name, class)

	result := database.DeleteUser(userId)
	if result != nil {
		ReturnModule.InternalServerError(w, "Error deleting member from leaderboard")
		return
	}
	ReturnModule.Success(w, r)
}
