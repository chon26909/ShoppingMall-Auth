package auth

import (
	"errors"
	"fmt"
)

type LoginWithEmailData struct {
	Email    string
	Password string
}

func (h *service) LoginWithEmail(data LoginWithEmailData) error {

	user, err := h.userRepository.GetUserByEmail(data.Email)
	if err != nil {
		return errors.New("user not found")
	}

	fmt.Println("user", user)

	return nil
}
