package auth

import "auth-service/repository"

type service struct {
	userRepository repository.UserRepository
}

type Service interface {
	LoginWithEmail(data LoginWithEmailData) error
}

func New(userRepository repository.UserRepository) Service {
	return &service{userRepository: userRepository}
}
