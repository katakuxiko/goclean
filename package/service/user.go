package service

import (
	"github.com/katakuxiko/clean_go/package/repository"
	"github.com/katakuxiko/clean_go/structure"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(userId int, userVariables structure.UsersVariables) (int, error) {
	return s.repo.Create(userId, userVariables)
}

func (s *UserService)  Update(userId int, input structure.UpdateUserVariables) error{
	// if err := input.Validate(); err != nil {
	// 	return err
	// }

	return s.repo.Update(userId, input)
}

func (s *UserService) GetAllVariables(userId int) (structure.UsersVariables,error){
	return s.repo.GetAllVariables(userId)
}

