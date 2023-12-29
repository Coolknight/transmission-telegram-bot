package main

import (
	"log"

	"github.com/Coolknight/transmission-telegram-bot/bot"
	"github.com/Coolknight/transmission-telegram-bot/config"
	"github.com/Coolknight/transmission-telegram-bot/transmission"
)

func main() {

	// Load configuration from config.yaml
	log.Println("Loading configuration")
	config, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// Initialize the Transmission client
	log.Println("Initialize transmission client")
	transmissionClient, err := transmission.NewClient(config.TransmissionURL, config.TransmissionUser, config.TransmissionPassword)
	if err != nil {
		log.Fatal("Error initializing Transmission client:", err)
		return
	}

	// Initialize the Telegram bot
	log.Println("Initialize Telegram Bot")
	telegramBot, err := bot.NewBot(config.BotToken)
	if err != nil {
		log.Fatal("Error initializing Telegram bot:", err)
	}

	// Handle incoming messages and commands for the bot
	telegramBot.Start(transmissionClient)

}
