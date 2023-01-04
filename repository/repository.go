package repository

import (
	"errors"

	"github.com/galihsatriawan/nested-call-mock/model"
)

var users []model.User = []model.User{
	{ID: 1},
	{ID: 2},
}

//go:generate go run github.com/golang/mock/mockgen --source=repository.go --package=mocks_repository --destination=mock/mock_repository.go
type Repository interface {
	IsUserExist(id int) bool
	GetUserByID(id int) (model.User, error)
}

type RepositoryImpl struct {
	users []model.User
}

func ProvideRepository() Repository {
	return &RepositoryImpl{
		users: users,
	}
}
func (repo *RepositoryImpl) IsUserExist(id int) bool {
	for _, user := range repo.users {
		if user.ID == id {
			return true
		}
	}
	return false
}
func (repo *RepositoryImpl) GetUserByID(id int) (model.User, error) {
	isExist := repo.IsUserExist(id)
	if !isExist {
		return model.User{}, errors.New("not found")
	}
	var userFound model.User
	for _, user := range repo.users {
		if user.ID == id {
			userFound = user
		}
	}
	return userFound, nil
}
