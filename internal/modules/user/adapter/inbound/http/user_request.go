package userhttp

type createUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
