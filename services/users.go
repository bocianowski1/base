package services

import (
	"github.com/KristianElde/butte/models"
	"github.com/KristianElde/butte/repo"
)

type IUserService interface {
	Create(email, password, name string) error
	FindByID(id string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(id, email, password, name string) error
	Delete(id string) error
}

type UserService struct {
	repo repo.IUserRepo
	// logger repo.ILogger
	// cache  repo.ICache
}

func NewUserService(repo repo.IUserRepo) IUserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(email, password, name string) error {
	user := &models.User{
		Email:    email,
		Password: password,
		Name:     name,
	}
	return s.repo.Create(user)
}

func (s *UserService) FindByID(id string) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) FindByEmail(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *UserService) Update(id, email, password, name string) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	user.Email = email
	user.Password = password
	user.Name = name

	return s.repo.Update(user)
}

func (s *UserService) Delete(id string) error {
	return s.repo.Delete(id)
}
