package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

func init() {
	Token = os.Getenv("DISCORD_TOKEN")
}

func main() {
	bot, err := discordgo.New("Bot " + Token)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	bot.AddHandler(messageCreate)

	err = bot.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	bot.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message is "ping" reply with "Pong!"
	if strings.ToLower(m.Content) == "king" {
		s.ChannelMessageSend(m.ChannelID, "Steve!")
	}

	if strings.ToLower(m.Content) == "your majesty" {
		s.ChannelMessageSendReply(m.ChannelID, "Yes?", m.Reference())
	}

	if strings.ToLower(m.Content) == "steve" {
		s.ChannelMessageSendReply(m.ChannelID, "That's my name, don't wear it out!", m.Reference())
	}

	if strings.ToLower(m.Content) == "!game" {
		game := generateRandomGame()
		s.ChannelMessageSendReply(m.ChannelID, game, m.Reference())
	}

}

func generateRandomGame() string {
	games := []string{"Brick Rigs", "Minecraft Java", "Forza Horizon 4", "Forza Horizon 5", "Minecraft Bedrock", "No Man's Sky", "Gang Beast", "Roblox"}

	random := rand.Intn(len(games))

	return games[random]
}
