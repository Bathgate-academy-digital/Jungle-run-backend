package structs

type ErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type SuccessResponse struct {
	Success bool `json:"success"`
}

type User struct {
	Name  string `json:"name"`
	Class string `json:"class"`
	Score int    `json:"score"`
}

var Users = []User{}
