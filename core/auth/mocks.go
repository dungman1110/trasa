package auth

import (
	"github.com/seknox/trasa/models"
	"github.com/stretchr/testify/mock"
)

type AuthMock struct {
	mock.Mock
}

func (a AuthMock) GetLoginDetails(trasaID, orgDomain string) (*models.UserWithPass, error) {
	args := a.Called(trasaID, orgDomain)
	return args.Get(0).(*models.UserWithPass), args.Error(1)

}

func (a AuthMock) Logout(sessionID string) error {
	return a.Called(sessionID).Error(0)
}
