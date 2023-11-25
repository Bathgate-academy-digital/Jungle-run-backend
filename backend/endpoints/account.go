package endpoints

import (
	"ba-digital/backend/endpoints/account"
	"ba-digital/backend/endpoints/account/creation"

	//"ba-digital/backend/endpoints/account/update"
	ReturnModule "ba-digital/backend/modules/return_module"
	"net/http"
)

func SubmitResult(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		account.PostRequest(w, r)
	} else {
		ReturnModule.MethodNotAllowed(w, r)
	}
}

func ManageUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//update.PostRequest(w, r)
	} else {
		ReturnModule.MethodNotAllowed(w, r)
	}
}

func ManageCreation(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		creation.PostRequest(w, r)
	} else {
		ReturnModule.MethodNotAllowed(w, r)
	}
}
