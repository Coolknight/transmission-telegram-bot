package main

import (
	"log"

	"github.com/Coolknight/transmission-telegram-bot/bot"
	"github.com/Coolknight/transmission-telegram-bot/config"
	"github.com/Coolknight/transmission-telegram-bot/solarman"
	"github.com/Coolknight/transmission-telegram-bot/transmission"
)

func main() {

	// Load configuration from config.yaml
	log.Println("Loading configuration")
	cfg, err := config.ReadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Initialize the Transmission client
	log.Println("Initialize transmission client")
	transmissionClient, err := transmission.NewClient(cfg.Transmission.URL, cfg.Transmission.User, cfg.Transmission.Password)
	if err != nil {
		log.Fatal("Error initializing Transmission client:", err)
		return
	}

	// Initialize the Telegram bot
	log.Println("Initialize Telegram Bot")
	telegramBot, err := bot.NewBot(cfg.Telegram.BotToken)
	if err != nil {
		log.Fatal("Error initializing Telegram bot:", err)
	}

	// Initialize solarman alerts daemon
	log.Println("Launch Solarman alert daemon")
	go solarman.ApiAlert(cfg)

	// Handle incoming messages and commands for the bot
	telegramBot.Start(transmissionClient)

}
