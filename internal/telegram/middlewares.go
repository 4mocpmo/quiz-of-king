package telegram

import (
	"context"
	"gopkg.in/telebot.v3"
	"quiz-of-kings/internal/entity"
)

func (t *Telegram) registerMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		acc := entity.Account{
			ID:          c.Sender().ID,
			FirstName:   c.Sender().FirstName,
			Username:    c.Sender().Username,
			DisplayName: c.Sender().Username,
		}
		account, created, err := t.App.Account.CreateOrUpdate(context.Background(), acc)
		if err != nil {
			return err
		}
		c.Set("account", account)
		c.Set("is_just_created", created)
		return next(c)
	}
}
