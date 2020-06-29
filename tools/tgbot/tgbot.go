package tgbot

import (
	"fmt"
	"net/http"
	"shop/models"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type ShopTgBot struct {
	tgBot  *tgbotapi.BotAPI
	chatId int64
}

func NewShopTgBot(token string, chatId int64) (*ShopTgBot, error) {
	cli := &http.Client{
		Timeout: 10 * time.Second,
	}
	bot, err := tgbotapi.NewBotAPIWithClient(token, cli)
	if err != nil {
		return nil, err
	}
	return &ShopTgBot{
		tgBot:  bot,
		chatId: chatId,
	}, nil
}

func (s *ShopTgBot) SendOrderNotification(order *models.Order) error {

	msgText := fmt.Sprintf("New order! \nPhone: %s, \nItemIDs: ", order.Phone)
	for _, itemID := range order.ItemIDs {
		msgText += fmt.Sprintf("\n\t%d", itemID)
	}

	msg := tgbotapi.NewMessage(s.chatId, msgText)

	_, err := s.tgBot.Send(msg)
	return err
}
