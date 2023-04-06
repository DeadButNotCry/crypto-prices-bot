package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgBot struct {
	bot      *tgbotapi.BotAPI
	adminIds []int
}

func NewTgBot(token string, adminIds []int) (*TgBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	if len(adminIds) == 0 {
		return nil, errors.New("admin list cannot be empty")
	}
	return &TgBot{
		bot:      bot,
		adminIds: adminIds,
	}, nil
}

func (b *TgBot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		switch {
		case update.Message != nil:
			go b.handleMessage(&update)
		case update.CallbackQuery != nil:
			go handleCallBackQuery(&update)

		}
	}

	return nil
}

