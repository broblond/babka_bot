package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/robfig/cron/v3"
	"gopkg.in/telebot.v3"
)

func calculateDaysSince(targetDate time.Time) int {
	currentDate := time.Now()
	diff := currentDate.Sub(targetDate)
	return int(diff.Hours() / 24)
}

func calculateDaysSinceSokr(sokrDate time.Time)int {
	currentDate := time.Now()
	diff := currentDate.Sub(sokrDate)
	return int(diff.Hours() / 24)
}

func sendMessage(bot *telebot.Bot, targetDate time.Time, chatID int64) {
	daysSince := calculateDaysSince(targetDate)
	message := fmt.Sprintf("Маша не выходит замуж %d дней", daysSince)

	bot.Send(&telebot.Chat{ID: chatID}, message)
}
func sendMessage(bot *telebot.Bot, sokrDate time.Time, chatID int64) {
	daysSince := calculateDaysSince(sokrDate)
	message := fmt.Sprintf("Дилару не сокращают %d дней", daysSince)

	bot.Send(&telebot.Chat{ID: chatID}, message)
}
func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")   // Токен  бота
	chatIDStr := os.Getenv("TELEGRAM_CHAT_ID") // ID чата, в который бот будет отправлять сообщения

	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		fmt.Println("Error parsing chat ID:", err)
		return
	}

	targetDate := time.Date(2020, time.January, 6, 0, 0, 0, 0, time.UTC)
	sokrDate := time.Date(2024, time.December, 19, 0, 0, 0, 0, time.UTC)

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		fmt.Println("Error creating bot:", err)
		return
	}

	bot.Handle("/skolko", func(c telebot.Context) error {
		daysSince := calculateDaysSince(targetDate)
		message := fmt.Sprintf("Маша не выходит замуж %d дней", daysSince)

		return c.Send(message)
	})
	
	bot.Handle("/sokr", func(c telebot.Context) error {
		daysSince := calculateDaysSinceSokr(sokrDate)
		message := fmt.Sprintf("Дилару не сокращают %d дней", daysSince)

		return c.Send(message)
	})

	c := cron.New(cron.WithLocation(time.FixedZone("MSK", 3*60*60))) // Московское время

	_, err = c.AddFunc("0 12 * * *", func() {
		sendMessage(bot, targetDate, chatID)
	})

	if err != nil {
		fmt.Println("Error setting up cron job:", err)
		return
	}

	c.Start()

	bot.Start()
}
