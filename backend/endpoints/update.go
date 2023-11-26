package endpoints

import (
	"ba-digital/backend/database"
	ReturnModule "ba-digital/backend/modules/return_module"
	"fmt"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.Form.Get("name")
		class := r.Form.Get("class")
		scoreStr := r.Form.Get("score")

		score, err := strconv.Atoi(scoreStr)
		if err != nil {
			fmt.Println(err)
			ReturnModule.CustomError(w, r, "Invalid score", 400)
			return
		}

		err = database.UpdateUser(name, class, score)
		if err != nil {
			ReturnModule.CustomError(w, r, "Failed to update user", 500)
			return
		}

		ReturnModule.Success(w, r)
	} else {
		ReturnModule.MethodNotAllowed(w, r)
	}
}
