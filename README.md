# Telegram Bot for Transmission Daemon

This Telegram bot was designed to interact with a Transmission daemon for initiating downloads using either torrent files or magnet links. But with the time I've found new exciting ways of using this bot so now is a multi-purpose bot :D

## Features

- Accepts commands:
  - `/torrent`: Upload a torrent file
  - `/magnet`: Input a magnet link
  - `/rss`: Adds a new feed to transmission-rss
  - `/screen`: This is a game for handling my kids screen time
  - `/scan`: Scans whatever is on the scanner tray and sends the scanned image back.
    - Possible subcommands are:
	  -  `/screen <kidname> start`
	  -  `/screen <kidname> add <minutes> <description>`
	  -  `/screen <kidname> take <minutes> <description>`
	  -  `/screen <kidname> log`
  - `/help`: Show available commands

Other than accepting commands it also serves as a SolarmanSmart API alert daemon so when the inverter is alerting it sends the alert through telegram.

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
transmission:
    url: "transmission_server_ip_address"
    user: "YOUR_TRANSMISSION_USERNAME"
    password: "YOUR_TRANSMISSION_PASSWORD"
solarman:
    appId: "YOUR_SOLARMAN_API_APPID"
    appSecret: "YOUR_SOLARMAN_API_APPSECRET"
    email: "YOUR_SOLARMAN_EMAIL_ACCOUNT"
    password: "YOUR_SHA256_ENCODED_PASSWORD"
api:
    authURL: https://globalapi.solarmanpv.com/account/v1.0/token
    apiURL: https://globalapi.solarmanpv.com/device/v1.0/currentData
telegram:
    botToken: "YOUR_TELEGRAM_BOT_TOKEN"
    chatID: "YOUR_TELEGRAM_CHATID"
device:
    deviceSn: "YOUR_DEVICE_SN"
```

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