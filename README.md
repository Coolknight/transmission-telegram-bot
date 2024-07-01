# Telegram Bot for Transmission Daemon

This Telegram bot was designed to interact with a Transmission daemon for initiating downloads using either torrent files or magnet links. However, it has evolved into a multi-purpose bot with new exciting features.

## Features

- Accepted commands:
  - `/torrent`: Upload a torrent file
  - `/magnet`: Input a magnet link
  - `/rss`: Adds a new feed to transmission-rss
  - `/scan`: Scans whatever is on the scanner tray and sends the scanned image back
  - `/screen`: This is a game for handling my kids screen time
    - Possible subcommands are:
	  -  `/screen <kidname> start`
	  -  `/screen <kidname> add <minutes> <description>`
	  -  `/screen <kidname> take <minutes> <description>`
	  -  `/screen <kidname> log`
  - `/help`: Show available commands

In addition to accepting commands, it also serves as a SolarmanSmart API alert daemon, sending alerts through Telegram when the inverter is alerting.

## Solarman Alerting Daemon
The Solarman Alerting Daemon is a crucial component of this Telegram bot. It enables real-time monitoring and alerting for SolarmanSmart API. By integrating with the Solarman API, the bot can send alerts through Telegram when the inverter is alerting. This feature ensures that users stay informed about any issues with their solar power system and can take prompt action.

To configure the Solarman Alerting Daemon, you need to provide your Solarman API credentials in the `config.yaml` file. Once configured, the bot will automatically monitor the inverter's status and send alerts whenever necessary.

Make sure to check the [Configuring the Bot](#configuring-the-bot) section for detailed instructions on setting up the Solarman Alerting Daemon.

Don't miss out on this powerful feature that enhances the functionality of the Telegram Bot for Transmission Daemon. Keep track of your solar power system effortlessly and receive timely notifications right in your Telegram chat.

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

- Start the bot by running the executable (`transmission-telegram-bot`).
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
