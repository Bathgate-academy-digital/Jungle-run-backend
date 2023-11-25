package structs

type ErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type HashedAndSaltedPassword struct {
	HashedPassword string
	RandomSalt     string
}

type SuccessResponse struct {
	Success bool `json:"success"`
}

type UserResponse struct {
	Username string `json:"username"`
	Ranking  int    `json:"ranking"`
}

var Users = []UserResponse{
	{Username: "Alex", Ranking: 3},
}
