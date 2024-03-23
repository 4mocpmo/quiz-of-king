package telegram

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
	"quiz-of-kings/internal/service"
	"time"
)

type Telegram struct {
	App *service.App
	bot *telebot.Bot
}

func NewTelegram(app *service.App, apiKey string) (*Telegram, error) {
	pref := telebot.Settings{
		Token:  apiKey,
		Poller: &telebot.LongPoller{Timeout: 50 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		logrus.WithError(err).Error("couldn't connect to telegram servers.")
		return nil, err
	}
	t := &Telegram{
		App: app, bot: b,
	}
	t.setupHandler()
	return t, nil

}

func (t *Telegram) Start() {
	t.bot.Start()
}
