# Telegram Bot for Transmission Daemon

This Telegram bot is designed to interact with a Transmission daemon for initiating downloads using either torrent files or magnet links.

## Features

- Accepts commands:
  - `/torrent`: Upload a torrent file
  - `/magnet`: Input a magnet link
  - `/rss`: Adds a new feed to transmission-rss
  - `/screen`: This is a game for handling my kids screen time
    - Possible subcommands are:
	  -  `/screen <kidname> start`
	  -  `/screen <kidname> add <minutes> <description>`
	  -  `/screen <kidname> take <minutes> <description>`
	  -  `/screen <kidname> log`
  - `/help`: Show available commands

## Installation and Setup

1. **Clone the Repository:**

    ```bash
    git clone https://github.com/Coolknight/transmission-telegram-bot.git
    ```

2. **Install Dependencies:**

    ```bash
    cd yourbot
    go mod tidy
    ```

3. **Configuration:**

    Create a `config.yaml` file in the `config` directory and fill in the required details (see [Configuring the Bot](#configuring-the-bot)).

4. **Build and Run:**

    ```bash
    go build -o transmission-telegram-bot main.go
    ./transmission-telegram-bot
    ```

## Configuring the Bot

Fill in the required details in the `config.yaml` file:

```yaml
bot_token: "YOUR_TELEGRAM_BOT_TOKEN"
transmission_url: "http://transmission_server_address:port/rpc"
transmission_user: "YOUR_TRANSMISSION_USERNAME"
transmission_password: "YOUR_TRANSMISSION_PASSWORD"
```

- bot_token: Your Telegram bot token obtained from BotFather.
- transmission_url: URL for your Transmission server RPC.
- transmission_user: Your Transmission username.
- transmission_password: Your Transmission password.

## Usage

- Start the bot by running the executable (transmission-telegram-bot).
- Interact with the bot via Telegram using the commands mentioned above.

### Docker-Compose

```yaml
version: '2'
services:
    telegram-bot:
        container_name: telegram-bot
        image: coolknight/transmission-telegram-bot:latest
        volumes:
        - <path to your config>:/config
        - <path to your transmission-rss>:/rss
        - /var/run/docker.sock:/var/run/docker.sock
        restart: unless-stopped
```

## Contributors

- [Manuel Mendoza](https://github.com/Coolknight)


## License

This project is licensed under the [MIT License](LICENSE).

## Contact or Support

For any inquiries or support, please contact [Manuel Mendoza](mailto:manumb@gmail.com).