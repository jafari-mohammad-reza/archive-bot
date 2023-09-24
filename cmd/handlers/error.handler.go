package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func ErrorHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, `Something wen wrong try again later`)
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}
