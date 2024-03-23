package integrationtest

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"quiz-of-kings/internal/entity"
	"quiz-of-kings/internal/repository"
	"quiz-of-kings/internal/repository/redis"
	"testing"
)

type testType struct {
	ID   string
	Name string
}

func (t testType) EntityId() entity.ID {
	return entity.NewID("testType", t.ID)
}

func TestCommonBehaviourSetAndGet(t *testing.T) {
	redisClient, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", redisPort))
	assert.NoError(t, err)
	fmt.Println("redis is now connected")

	ctx := context.Background()
	cb := repository.NewRedisCommandBehaviour[testType](redisClient)

	//test save and get
	err = cb.Save(ctx, &testType{ID: "12", Name: "mostafa"})
	assert.NoError(t, err)
	val, err := cb.Get(ctx, entity.NewID("testType", "12"))
	assert.NoError(t, err)
	assert.Equal(t, "mostafa", val.Name)
	assert.Equal(t, "12", val.ID)

	//test update
	err = cb.Save(ctx, &testType{ID: "12", Name: "Ali"})
	assert.NoError(t, err)
	val, err = cb.Get(ctx, entity.NewID("testType", "12"))
	assert.NoError(t, err)
	assert.Equal(t, "Ali", val.Name)
	assert.Equal(t, "12", val.ID)

	val, err = cb.Get(ctx, entity.NewID("testType", "234"))
	assert.ErrorIs(t, repository.ErrorNotFound, err)

	redisClient.Close()
}
