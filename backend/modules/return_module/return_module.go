package return_module

import (
	"encoding/json"
	"jungle-rush/backend/structs"
	"log"
	"net/http"
)

func respond(w http.ResponseWriter, r *http.Request, jsonObject any, statusCode int) {
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

func CustomError(w http.ResponseWriter, r *http.Request, errorMessage string, errorCode int) {
	errorResponse := structs.ErrorResponse{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	}
	respond(w, r, errorResponse, errorCode)
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	errorMessage := "That method is not accepted at this endpoint."
	CustomError(w, r, errorMessage, http.StatusMethodNotAllowed)
}

func BadRequest(w http.ResponseWriter, r *http.Request, errorMessage string) {
	CustomError(w, r, errorMessage, http.StatusBadRequest)
}

func InternalServerError(w http.ResponseWriter, r *http.Request, errorMessage string) {
	CustomError(w, r, errorMessage, http.StatusInternalServerError)
}

func ID(w http.ResponseWriter, r *http.Request, id int) {
	idResponse := structs.IdResponse{Id: id}
	respond(w, r, idResponse, http.StatusOK)
}

func Leaderboard(w http.ResponseWriter, r *http.Request, leaderboard []structs.User) {
	respond(w, r, leaderboard, http.StatusOK)
}

func Success(w http.ResponseWriter, r *http.Request) {
	successResponse := structs.SuccessResponse{Success: true}
	respond(w, r, successResponse, http.StatusOK)
}
