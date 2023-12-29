package bot

import (
	"fmt"
	"log"
	"time"

	"github.com/Coolknight/transmission-telegram-bot/transmission"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Bot struct holds the Telegram bot
type Bot struct {
	BotAPI *tgbotapi.BotAPI
}

// NewBot initializes a new Telegram bot
func NewBot(token string) (*Bot, error) {
	botAPI, err := tgbotapi.NewBotAPI(token)
	fmt.Println("NewBot")
	if err != nil {
		return nil, err
	}
	return &Bot{BotAPI: botAPI}, nil
}

// Start updates handler for the bot
func (b *Bot) Start(transmission *transmission.Client) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.BotAPI.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Text {
		case "/torrent":
			b.HandleDownloadCommand(updates, update.Message.Chat.ID, transmission, true)
		case "/magnet":
			b.HandleDownloadCommand(updates, update.Message.Chat.ID, transmission, false)
		case "/help":
			b.HandleHelpCommand(update)
		default:
			b.HandleDefault(update)
		}
	}
	fmt.Println("d")
}

// HandleDownloadCommand handles /torrent and /magnet commands
func (b *Bot) HandleDownloadCommand(updates <-chan tgbotapi.Update, chatID int64, transmission *transmission.Client, isTorrent bool) {
	// Ask for either a torrent file or magnet link based on isTorrent flag
	var requestMessage string
	if isTorrent {
		requestMessage = "Please upload the torrent file."
	} else {
		requestMessage = "Please enter the magnet link:"
	}

	msg := tgbotapi.NewMessage(chatID, requestMessage)
	b.BotAPI.Send(msg)

	// Initialize the variables for file/link and download path
	var fileLink, downloadPath string

	// Listen for the user's input
	for update := range updates {
		var err error
		if update.Message == nil {
			continue
		}
		// Handle the user's input based on isTorrent flag
		if isTorrent {
			if update.Message.Document != nil {
				fileLink, err = b.BotAPI.GetFileDirectURL(update.Message.Document.FileID)
				if err != nil {
					log.Println("Error getting file link:", err)
					// Handle the error as needed
					return
				}
			}
			break
		} else {
			fileLink = update.Message.Text
			break
		}
	}

	// Ask for the download path
	msg = tgbotapi.NewMessage(chatID, "Enter the download path:")
	b.BotAPI.Send(msg)

	// Listen for the user's input for the download path
	for update := range updates {
		if update.Message == nil {
			continue
		}
		// Extract the download path
		downloadPath = update.Message.Text
		break
	}

	// Start the download using the provided file/link and download path via the Transmission client
	torrentID, err := transmission.StartDownload(fileLink, downloadPath)
	if err != nil {
		log.Println("Error starting download:", err)
		return
	}

	// Notify the user that the download has started
	startMsg := tgbotapi.NewMessage(chatID, "Download started!")
	b.BotAPI.Send(startMsg)

	// Poll download status every minute until it's completed
	go b.WaitForDownload(torrentID, chatID, transmission)
}

// WaitForDownload is designed to be launched as a subroutine and wait for the download and inform the user
func (b *Bot) WaitForDownload(torrentID, chatID int64, transmission *transmission.Client) error {
	for {
		time.Sleep(time.Minute)

		// Check if download is complete
		isComplete, err := transmission.IsDownloadComplete(torrentID)
		if err != nil {
			log.Println("Error checking download status:", err)
			return err
		}

		if isComplete {
			// If download is complete, get the torrent name
			name, err := transmission.GetName(torrentID)
			if err != nil {
				log.Println("Error retrieving torrent info:", err)
				return err
			}

			// Notify the user about the completed download with the torrent name
			completedMsg := tgbotapi.NewMessage(chatID, "Download completed: "+name)
			b.BotAPI.Send(completedMsg)
			break // Exit the loop when download is complete
		}
	}
	return nil
}

// HandleHelpCommand handles the /help command
func (b *Bot) HandleHelpCommand(update tgbotapi.Update) {
	// Get the chat ID
	chatID := update.Message.Chat.ID

	// Create the help message with available commands
	helpMessage := "Available commands:\n" +
		"/torrent - Upload a torrent file\n" +
		"/magnet - Input a magnet link\n" +
		"/help - Show available commands"

	// Send the help message to the user
	msg := tgbotapi.NewMessage(chatID, helpMessage)
	_, err := b.BotAPI.Send(msg)
	if err != nil {
		log.Println("Error sending help message:", err)
	}
}

// HandleDefault handles any unrecognized command or input
func (b *Bot) HandleDefault(update tgbotapi.Update) {
	// Get the chat ID
	chatID := update.Message.Chat.ID

	// Send an error message for unrecognized commands
	errorMessage := "Sorry, I don't recognize that command. Please use /help to see available commands."
	msg := tgbotapi.NewMessage(chatID, errorMessage)
	_, err := b.BotAPI.Send(msg)
	if err != nil {
		log.Println("Error sending default message:", err)
	}
}
