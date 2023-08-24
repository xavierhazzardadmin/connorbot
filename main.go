package main

import (
	"fmt"
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

	if strings.ToLower(m.Content) == "Your Majesty" {
		s.ChannelMessageSendReply(m.ChannelID, "Yes?", m.Reference())
	}

	if strings.ToLower(m.Content) == "Steve" {
		s.ChannelMessageSendReply(m.ChannelID, "That's my name, don't wear it out!", m.Reference())
	}

	// If the message is "pong" reply with "Ping!"
	// if strings.ToLower(m.Content) == "pong" {
	// 	s.ChannelMessageSend(m.ChannelID, "Ping!")
	// }
	//
	// if strings.ToLower(m.Content) == "ping pong" {
	// 	s.ChannelMessageSend(m.ChannelID, "Pong! Ping!")
	// }
}
