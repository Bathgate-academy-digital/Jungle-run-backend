package return_module

import (
	"encoding/json"
	"jungle-rush/backend/structs"
	"log"
	"net/http"
)

func respond(w http.ResponseWriter, jsonObject any, statusCode int) {
	responseContent, err := json.Marshal(jsonObject)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(statusCode)
	_, err = w.Write(responseContent)
	if err != nil {
		log.Fatal(err)
	}
}

func CustomError(w http.ResponseWriter, errorMessage string, errorCode int) {
	errorResponse := structs.ErrorResponse{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	}
	respond(w, errorResponse, errorCode)
}

func MethodNotAllowed(w http.ResponseWriter) {
	errorMessage := "That method is not accepted at this endpoint."
	CustomError(w, errorMessage, http.StatusMethodNotAllowed)
}

func BadRequest(w http.ResponseWriter, errorMessage string) {
	CustomError(w, errorMessage, http.StatusBadRequest)
}

func InternalServerError(w http.ResponseWriter, errorMessage string) {
	CustomError(w, errorMessage, http.StatusInternalServerError)
}

func ID(w http.ResponseWriter, id int) {
	idResponse := structs.IdResponse{Id: id}
	respond(w, idResponse, http.StatusOK)
}

func Leaderboard(w http.ResponseWriter, leaderboard []structs.User) {
	respond(w, leaderboard, http.StatusOK)
}

func Success(w http.ResponseWriter, r *http.Request) {
	successResponse := structs.SuccessResponse{Success: true}
	respond(w, successResponse, http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, user structs.User) {
	respond(w, user, http.StatusOK)
}
