package account

import (
	"ba-digital/backend/database"
	ReturnModule "ba-digital/backend/modules/return_module"
	"ba-digital/backend/structs"
	"net/http"
)

func GetRequest(w http.ResponseWriter, r *http.Request) {
	sessionAuthentication := r.Header.Get("session")
	username := r.Header.Get("username")

	account := structs.UserResponse{
		Username: "alex",
		Ranking:  4,
	}
	ReturnModule.AccountData(w, r, account)
	return
	if sessionAuthentication != "" {
		accountDataStruct := database.GetAccountDataFromSession(sessionAuthentication)

		if accountDataStruct.Username == "" {
			ReturnModule.InternalServerError(w, r)
		} else {
			ReturnModule.AccountData(w, r, accountDataStruct)
		}
	} else if username != "" {
		//accountDataStruct := database.GetAccountData(username)
		//
		//if accountDataStruct.Username == "" {
		//	ReturnModule.InternalServerError(w, r)
		//} else {
		//	ReturnModule.AccountData(w, r, accountDataStruct)
		//}
	} else {
		ReturnModule.MissingData(w, r)
	}
}
