package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/telebot.v3"
)

func daysSince() int {
	startDate := time.Date(2020, time.January, 6, 0, 0, 0, 0, time.UTC)
	now := time.Now()
	duration := now.Sub(startDate)
	return int(duration.Hours() / 24)
}

func sendMessage(b *telebot.Bot, chatID int64) {
	days := daysSince()
	message := fmt.Sprintf("Маша не выходит замужем %d дней", days)
	_, err := b.Send(telebot.ChatID(chatID), message)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

func waitUntilNoon() time.Duration {
	now := time.Now()
	nextNoon := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location())
	if now.After(nextNoon) {
		nextNoon = nextNoon.Add(24 * time.Hour)
	}
	return time.Until(nextNoon)
}

func main() {
	token := "YOUR_TELEGRAM_BOT_API_TOKEN"
	chatID := int64(YOUR_CHAT_ID)

	b, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	for {
		time.Sleep(waitUntilNoon())
		sendMessage(b, chatID)
	}
}
