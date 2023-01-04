package service

import (
	"github.com/galihsatriawan/nested-call-mock/model"
	"github.com/galihsatriawan/nested-call-mock/repository"
)

//go:generate go run github.com/golang/mock/mockgen --source=service.go --package=mocks_service --destination=mock/mock_service.go
type Service interface {
	GetUserByID(id int) (model.User, error)
	Eligible(id int) bool
}
type Option func(serviceImpl *ServiceImpl)

func WithThisMock(this Service) Option {
	return func(serviceImpl *ServiceImpl) {
		serviceImpl.this = this
	}
}

type ServiceImpl struct {
	this Service
	repo repository.Repository
}

func ProvideService(repo repository.Repository, opts ...Option) Service {
	svc := &ServiceImpl{
		repo: repo,
	}
	for _, opt := range opts {
		opt(svc)
	}
	if svc.this == nil {
		svc.this = svc
	}
	return svc
}
func (s *ServiceImpl) GetUserByID(id int) (model.User, error) {
	isEligible := s.this.Eligible(id)
	if !isEligible {
		return model.User{}, nil
	}
	return s.repo.GetUserByID(id)
}

func (s *ServiceImpl) Eligible(id int) bool {
	return s.repo.IsUserExist(id)
}
