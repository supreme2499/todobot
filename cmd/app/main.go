package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"telegrambottodo/internal/config"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("download error .env file: %v", err)
	}
	cfg := config.MustLoad()
	fmt.Println("cfg compleate")
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Authorized on account %s", bot.Self.UserName)

}
