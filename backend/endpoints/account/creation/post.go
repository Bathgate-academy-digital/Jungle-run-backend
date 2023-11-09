package creation

import (
	ReturnModule "ba-digital/backend/modules/return_module"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func PostRequest(w http.ResponseWriter, r *http.Request) {

	username, err := generateUsername(r)
	if err != nil {
		ReturnModule.CustomError(w, r, "invalid data format", http.StatusBadRequest)
		return
	}
	fmt.Println(username)
}

func generateUsername(req *http.Request) (string, error) {
	var body map[string]interface{}
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		return "", err
	}

	firstName, firstNameOk := body["first_name"].(string)
	lastName, lastNameOk := body["last_name"].(string)
	classNumber, classNumberOk := body["class_number"].(string)

	if !firstNameOk || !lastNameOk || !classNumberOk {
		return "", fmt.Errorf("invalid data format")
	}

	username := fmt.Sprintf("%s%s%s", strings.ToLower(firstName), strings.ToLower(lastName), classNumber)
	username = strings.ReplaceAll(username, " ", "")
	return username, nil
}
