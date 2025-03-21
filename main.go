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

func calculateDaysTill(targetDate time.Time) int {
	currentDate := time.Now()
	if targetDate > currentDate {
		diff := targetDate.Sub(currentDate)
		return
	}
	else {
		diff := currentDate.Sub(targetDate)
		return
	}	
	return int(diff.Hours() / 24)
}

func sendMessage(bot *telebot.Bot, targetDate time.Time, user_type string, messageText string, chatID int64) {
	if user_type = "Masha" {
		daysSince := calculateDaysSince(targetDate)
		return
	}
	else {
		daysSince := calculateDaysTill(targetDate)
		return
	}
	
	message := fmt.Sprintf(messageText, daysSince)

	bot.Send(&telebot.Chat{ID: chatID}, message)
}

func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")   
	chatIDStr := os.Getenv("TELEGRAM_CHAT_ID") 
	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		fmt.Println("Error parsing chat ID:", err)
		return
	}

	targetDate := time.Date(2020, time.January, 6, 0, 0, 0, 0, time.UTC) 
	probationDate := time.Date(2025, time.June, 10, 0, 0, 0, 0, time.UTC) 

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
		daysSince := calculateDaysTill(probationDate)
		message := fmt.Sprintf("Диларе до конца испыталки %d дней", daysSince)
		return c.Send(message)
	})

	c := cron.New(cron.WithLocation(time.FixedZone("MSK", 3*60*60))) 
	_, err = c.AddFunc("0 12 * * *", func() {
		sendMessage(bot, targetDate, "Masha", "Маша не выходит замуж %d дней", chatID)
		
		currentDate := time.Now()
		if probationDate > currentDate {
			sendMessage(bot, probationDate, "Dilara", "Диларе до конца испыталки %d дней", chatID)
			return
		}
		else {
			sendMessage(bot, probationDate, "Dilara", "Дилара закрыла испыталку %d дней назад", chatID)
			return
		}		
	})

	if err != nil {
		fmt.Println("Error setting up cron job:", err)
		return
	}

	c.Start()

	bot.Start()
}
