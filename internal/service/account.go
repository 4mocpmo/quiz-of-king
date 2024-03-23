package service

import (
	"context"
	"errors"
	"quiz-of-kings/internal/entity"
	"quiz-of-kings/internal/repository"
	"time"
)

var (
	DefaultState = "home"
)

type AccountService struct {
	accounts repository.AccountRepository
}

func NewAccountService(rep repository.AccountRepository) *AccountService {
	return &AccountService{accounts: rep}
}

// CreateOrUpdate creates a new user in the datasource or update existing user
func (a *AccountService) CreateOrUpdate(ctx context.Context, account entity.Account) (entity.Account, bool, error) {
	savedAccount, err := a.accounts.Get(ctx, account.EntityId())
	if err == nil {
		if savedAccount.Username != account.Username || savedAccount.FirstName != account.FirstName {
			savedAccount.Username = account.Username
			savedAccount.FirstName = account.FirstName
			return savedAccount, false, a.accounts.Save(ctx, savedAccount)
		}
		return savedAccount, false, nil
	}
	//user does not exists in the database
	if errors.Is(err, repository.ErrorNotFound) {
		account.JoinedAt = time.Now()
		account.State = DefaultState
		return account, true, a.accounts.Save(ctx, account)

	}
	return entity.Account{}, false, err
}
