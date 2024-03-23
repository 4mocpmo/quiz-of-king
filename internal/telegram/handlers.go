package telegram

import "gopkg.in/telebot.v3"

func (t *Telegram) setupHandler() {

	//middleware
	t.bot.Use(t.registerMiddleware)

	//handlers
	t.bot.Handle("/start", t.start)
}

func (t *Telegram) start(c telebot.Context) error {
	return c.Reply("Hello World")
}
