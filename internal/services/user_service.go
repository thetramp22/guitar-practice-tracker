package services

import (
	"github.com/thetramp22/rifflog/internal/models"
	"github.com/thetramp22/rifflog/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) RegisterUser(req models.RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := models.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	return s.Repo.CreateUser(user)
}
