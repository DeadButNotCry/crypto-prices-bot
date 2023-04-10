package telegram

import (
	"log"

	"github.com/deadbutnotcry/crypto-prices-bot/internal/consts/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleCallBackQuery(u *tgbotapi.Update) {

}

func (b *TgBot) handleMessage(u *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(u.Message.From.ID, messages.RefLinkRU)
	msg.ParseMode = "html"
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
}
