package service_test

import (
	"testing"

	"github.com/galihsatriawan/nested-call-mock/model"
	mockRepo "github.com/galihsatriawan/nested-call-mock/repository/mock"
	svc "github.com/galihsatriawan/nested-call-mock/service"
	mockSvc "github.com/galihsatriawan/nested-call-mock/service/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMocks := mockRepo.NewMockRepository(ctrl)
	svcMocks := mockSvc.NewMockService(ctrl)
	id := 12
	expectUser := model.User{ID: id}

	svcMocks.EXPECT().Eligible(id).Return(true)
	repoMocks.EXPECT().GetUserByID(id).
		Return(expectUser, nil)

	s := svc.ProvideService(repoMocks, svc.WithThisMock(svcMocks))
	user, err := s.GetUserByID(id)

	assert.Equal(t, user, expectUser)
	assert.Equal(t, err, nil)
}
