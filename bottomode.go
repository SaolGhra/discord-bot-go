package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

// Config represents the configuration file structure
type Config struct {
	Token string `json:"token"`
}

var config Config

func main() {
	// Load the configuration
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}

	// Create a new Discord session
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	// Register the messageCreate func as a callback for MessageCreate events
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening connection: %v", err)
	}

	// Wait here until CTRL-C or other term signal is received
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	select {}
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message is "/meme", fetch and post a meme
	if m.Content == "/meme" {
		meme, err := getRandomMeme()
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Failed to fetch a meme.")
			return
		}
		s.ChannelMessageSend(m.ChannelID, meme)
	}
}

// getRandomMeme fetches a random meme from Reddit
func getRandomMeme() (string, error) {
	url := "https://www.reddit.com/r/memes/random/.json"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result []interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	if len(result) == 0 {
		return "", fmt.Errorf("no meme found")
	}

	postData := result[0].(map[string]interface{})
	data := postData["data"].(map[string]interface{})
	children := data["children"].([]interface{})
	if len(children) == 0 {
		return "", fmt.Errorf("no meme found")
	}

	memeData := children[0].(map[string]interface{})["data"].(map[string]interface{})
	title := memeData["title"].(string)
	url = memeData["url"].(string)

	return fmt.Sprintf("%s\n%s", title, url), nil
}
