package repository

import (
	"context"
	"errors"
	"github.com/redis/rueidis"
	"github.com/sirupsen/logrus"
	"quiz-of-kings/internal/entity"
	"quiz-of-kings/pkg/jsonhelper"
)

var _ CommonBehaviour[entity.Entity] = &RedisCommandBehaviour[entity.Entity]{}

type RedisCommandBehaviour[T entity.Entity] struct {
	client rueidis.Client
}

func NewRedisCommandBehaviour[T entity.Entity](client rueidis.Client) *RedisCommandBehaviour[T] {
	return &RedisCommandBehaviour[T]{
		client: client,
	}
}

func (r RedisCommandBehaviour[T]) Get(ctx context.Context, id entity.ID) (T, error) {
	var t T
	cmd := r.client.B().JsonGet().Key(id.String()).Path(".").Build()
	val, err := r.client.Do(ctx, cmd).ToString()
	if err != nil {
		//handle redis nil error
		if errors.Is(err, rueidis.Nil) {
			return t, ErrorNotFound
		}
		logrus.WithError(err).WithField("id", id).Errorln("couldn't retrieve from Redis")
		return t, err
	}
	return jsonhelper.Decode[T]([]byte(val)), nil
}

func (r RedisCommandBehaviour[T]) Save(ctx context.Context, entity entity.Entity) error {
	cmd := r.client.B().JsonSet().Key(entity.EntityId().String()).Path("$").Value(string(jsonhelper.Encode(entity))).Build()
	if err := r.client.Do(ctx, cmd).Error(); err != nil {
		logrus.WithError(err).WithField("entity", entity).Errorln("couldn't save the entity")
		return err
	}
	return nil
}
