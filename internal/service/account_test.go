package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"quiz-of-kings/internal/entity"
	"quiz-of-kings/internal/repository"
	"quiz-of-kings/internal/repository/mocks"
	"testing"
)

func TestAccountService_CreateOrUpdate_WhileUserNotExists(t *testing.T) {
	accRep := &mocks.AccountRepository{}
	service := NewAccountService(accRep)

	accRep.On("Get", mock.Anything, entity.NewID("account", 12)).
		Return(
			entity.Account{}, repository.ErrorNotFound,
		).Once()

	accRep.On("Save", mock.Anything, mock.MatchedBy(func(acc entity.Account) bool {
		return acc.FirstName == "mostafa"
	})).Return(nil).Once()

	newAcc, created, err := service.CreateOrUpdate(context.Background(), entity.Account{ID: 12, FirstName: "mostafa"})
	assert.NoError(t, err)
	assert.Equal(t, true, created)
	assert.Equal(t, "mostafa", newAcc.FirstName)

	accRep.AssertExpectations(t)
}

func TestAccountService_CreateOrUpdate_WhileUserExists(t *testing.T) {
	accRep := &mocks.AccountRepository{}
	service := NewAccountService(accRep)

	accRep.On("Get", mock.Anything, entity.NewID("account", 12)).
		Return(
			entity.Account{ID: 12, FirstName: "Reza"}, nil,
		).Once()

	accRep.On("Save", mock.Anything, mock.MatchedBy(func(acc entity.Account) bool {
		return acc.FirstName == "mostafa"
	})).Return(nil).Once()

	newAcc, created, err := service.CreateOrUpdate(context.Background(), entity.Account{ID: 12, FirstName: "mostafa"})
	assert.NoError(t, err)
	assert.Equal(t, false, created)
	assert.Equal(t, "mostafa", newAcc.FirstName)

	accRep.AssertExpectations(t)
}

func TestAccountService_CreateOrUpdate_WhileUserHasNotChange(t *testing.T) {
	accRep := &mocks.AccountRepository{}
	service := NewAccountService(accRep)

	accRep.On("Get", mock.Anything, entity.NewID("account", 12)).
		Return(
			entity.Account{ID: 12, FirstName: "mostafa"}, nil,
		).Once()

	newAcc, created, err := service.CreateOrUpdate(context.Background(), entity.Account{ID: 12, FirstName: "mostafa"})
	assert.NoError(t, err)
	assert.Equal(t, false, created)
	assert.Equal(t, "mostafa", newAcc.FirstName)

	accRep.AssertExpectations(t)
}
