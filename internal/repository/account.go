package repository

import (
	"github.com/redis/rueidis"
	"quiz-of-kings/internal/entity"
)

var _ AccountRepository = &AccountRedisRepository{}

type AccountRedisRepository struct {
	*RedisCommandBehaviour[entity.Account]
}

func NewAccountRedisRepository(client rueidis.Client) *AccountRedisRepository {
	return &AccountRedisRepository{
		NewRedisCommandBehaviour[entity.Account](client),
	}
}
