package auth

type UserAuthenticationResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
