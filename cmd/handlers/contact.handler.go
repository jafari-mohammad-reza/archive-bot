package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// func /contact - *Provide the contact informations to user*.
func ContactHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, `
Contact me by [email](mailto:mohammadrexzajafari.dev@gmail.com)
on Telegram: [DaYeezus](https://t.me/DaYeezus)
[Github Repo](https://github.com/jafari-mohammad-reza/archive-bot)
`)
	msg.ParseMode = "Markdown"
	_, err := bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
