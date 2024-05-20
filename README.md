# discord-bot-go

![GitHub Repo stars](https://img.shields.io/github/stars/saolghra/discord-bot-go?style=for-the-badge)

This is a Discord bot written in Golang that scrapes Reddit for random memes.

## Features

- **/meme command**: Use the `/meme` command in your Discord server to get a random meme from Reddit.
- **Golang**: The bot is written in Golang, a powerful and efficient programming language.

## Prerequisites

Before running the bot, make sure you have the following:

- **Golang**: Install Golang on your machine. You can download it from the official Golang website.
- **Discord Bot Token**: Obtain a Discord bot token by creating a new bot on the Discord Developer Portal.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/discord-bot-go.git
   ```

2. Navigate to the project directory:

   ```bash
   cd discord-bot-go
   ```

3. Install the required dependencies:

   ```bash
   go get
   ```

4. Set your Discord bot token:

   - Rename the `config-example.json` file to `config.json`.
   - Open the `config.json` file and replace the placeholder token with your actual Discord bot token.
   - Save the `config.json` file.

5. Run the bot:

   ```bash
   go run main.go
   ```

## Usage

Once the bot is running and connected to your Discord server, you can use the following command:

- **/meme**: Use this command to get a random meme from Reddit.

## Contributing

Contributions are welcome! If you have any suggestions, bug reports, or feature requests, please open an issue or submit a pull request.
