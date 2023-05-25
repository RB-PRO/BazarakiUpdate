package app

import (
	"context"
	"strconv"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
)

type Notification struct {
	*telegram.Telegram
}

func NewNotification(token, ChatID string) (Notification, error) {

	// Create a telegram service. Ignoring error for demo simplicity.
	telegramService, ErrorServece := telegram.New(token)
	if ErrorServece != nil {
		return Notification{}, ErrorServece
	}

	// Переводить ChatID из string в int64
	ChatID_int, ErrParseInt := strconv.ParseInt(ChatID, 10, 64)
	if ErrParseInt != nil {
		return Notification{}, ErrParseInt
	}

	// Добавить ID,куда будут посылаться уведомления
	telegramService.AddReceivers(ChatID_int)

	// Tell our notifier to use the telegram service. You can repeat the above process
	// for as many services as you like and just tell the notifier to use them.
	// Inspired by http middlewares used in higher level libraries.
	notify.UseServices(telegramService)

	// Send a test message.
	ErrorTelegramSend := notify.Send(
		context.Background(),
		"Отчёт о работе",
		"BazarakiUpdate: Начинаю работу",
	)
	if ErrorTelegramSend != nil {
		return Notification{}, ErrorTelegramSend
	}

	return Notification{telegramService}, nil
}

func (notif Notification) Sends(message string) error {
	return notif.Send(
		context.Background(),
		"Отчёт о работе",
		message,
	)
}
