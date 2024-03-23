package integrationtest

import (
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/sirupsen/logrus"
	"os"
	"quiz-of-kings/internal/repository/redis"
	"quiz-of-kings/pkg/testhelper"
	"testing"
)

var redisPort string

func TestMain(m *testing.M) {

	if !testhelper.IsIntegration() {
		fmt.Println("what the fuck?")
		return
	}

	pool := testhelper.StartDockerPull()

	//set up the redis container for tests
	redisResource := testhelper.StartDockerInstance(pool, "redis/redis-stack-server", "latest",
		func(res *dockertest.Resource) error {
			port := res.GetPort("6379/tcp")
			_, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", port))
			return err
		})
	redisPort = redisResource.GetPort("6379/tcp")

	fmt.Printf("redis is up and running on port %s\n", redisPort)
	exitCode := m.Run()
	err := redisResource.Close()
	if err != nil {
		logrus.WithError(err).Fatal("couldn't close the redis resource")
	}
	os.Exit(exitCode)
}
