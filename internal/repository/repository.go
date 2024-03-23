package repository

import (
	"context"
	"errors"
	"quiz-of-kings/internal/entity"
)

var (
	ErrorNotFound = errors.New("entity not found")
)

type CommonBehaviour[T entity.Entity] interface {
	Get(ctx context.Context, id entity.ID) (T, error)
	Save(ctx context.Context, entity entity.Entity) error
}

type AccountRepository interface {
	CommonBehaviour[entity.Account]
}
