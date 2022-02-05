package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/suhrr/rakuten-card-statement-bot/message"
	"github.com/suhrr/rakuten-card-statement-bot/rakuten"
)

var env = map[string]string{
	"RAKUTEN_ID":          "",
	"RAKUTEN_PASSWORD":    "",
	"RAKUTEN_CARD_NAME":   "",
	"LINE_CHANNEL_SECRET": "",
	"LINE_CHANNEL_TOKEN":  "",
}

func main() {
	fmt.Println("Start")

	if err := loadEnv(); err != nil {
		log.Fatal(err)
	}

	cookies, err := rakuten.GetLoggedInCookies(env["RAKUTEN_ID"], env["RAKUTEN_PASSWORD"], env["RAKUTEN_CARD_NAME"])
	if err != nil {
		log.Fatal(err)
	}

	records, err := rakuten.ReadStatementCsv(cookies)
	if err != nil {
		log.Fatal(err)
	}

	s, err := rakuten.NewStatement(records)
	if err != nil {
		log.Fatal(err)
	}

	bot, err := linebot.New(env["LINE_CHANNEL_SECRET"], env["LINE_CHANNEL_TOKEN"])
	if err != nil {
		log.Fatal(err)
	}

	m := message.NewMonthlyTotalMessage(rakuten.GetMonthlyTotal(s))
	if _, err := bot.BroadcastMessage(m).Do(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Send messages.")
}

func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	for index := range env {
		_env := os.Getenv(index)
		if _env == "" {
			return err
		}
		env[index] = _env
	}

	return nil
}
