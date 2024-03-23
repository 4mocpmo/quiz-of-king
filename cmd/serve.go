package cmd

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"quiz-of-kings/internal/repository"
	"quiz-of-kings/internal/repository/redis"
	"quiz-of-kings/internal/service"
	"quiz-of-kings/internal/telegram"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the telegram bot",
	Run:   serve,
}

func serve(_ *cobra.Command, _ []string) {
	_ = godotenv.Load()
	//connect to repositories
	redisClient, err := redis.NewRedisClient(os.Getenv("REDIS_URL"))
	if err != nil {
		logrus.WithError(err).Fatalln("couldn't connect to the redis server")
	}
	redisRepository := repository.NewAccountRedisRepository(redisClient)

	//set up app
	app := service.NewApp(service.NewAccountService(redisRepository))
	tg, err := telegram.NewTelegram(app, os.Getenv("BOT_API"))
	if err != nil {
		logrus.WithError(err).Error("couldn't connect to the telegram server")
	}
	tg.Start()

}

func init() {
	rootCmd.AddCommand(serveCmd)
}
