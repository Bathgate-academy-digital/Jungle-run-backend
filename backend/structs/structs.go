package structs

type ErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type SuccessResponse struct {
	Success bool `json:"success"`
}

type IdResponse struct {
	Id int `json:"id"`
}

type User struct {
	Name  string `json:"name"`
	Class string `json:"class"`
	Time  string `json:"time"`
}
