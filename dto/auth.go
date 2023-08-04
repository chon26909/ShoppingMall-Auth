package dto

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	// Firstname string `json:"firstname"`
	// Lastname  string `json:"lastname"`
}
