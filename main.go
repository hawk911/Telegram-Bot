package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/telegram-bot-api.v4"
)

var (
	telegramBotToken string
	numericKeyboard  = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("1"),
			tgbotapi.NewKeyboardButton("2"),
			tgbotapi.NewKeyboardButton("3"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("4"),
			tgbotapi.NewKeyboardButton("5"),
			tgbotapi.NewKeyboardButton("6"),
		),
	)
)

func init() {
	flag.StringVar(&telegramBotToken, "telegramtokenbotapi", "", "Telefram Token Api")
	flag.Parse()

	if telegramBotToken == "" {
		log.Print("-tokenbotapi is empty!")
		os.Exit(1)

	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		//bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Hello %s. I am a bot! Write /help if you don't know who I", update.Message.From.UserName)))
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "help":
				msg.Text = "type /say or /status."
			case "say":
				msg.Text = "Hi :)"
			case "status":
				msg.Text = "I'm ok."
			case "start":
				msg.Text = fmt.Sprintf("Hello %s. I am a bot! Write /help if you don't know who I", update.Message.From.UserName)
			default:
				msg.Text = "I don't know that command"
			}
		} else {
			switch update.Message.Text {
			case "open":
				msg.ReplyMarkup = numericKeyboard
			case "close":
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			}
		}
		bot.Send(msg)
	}
}
