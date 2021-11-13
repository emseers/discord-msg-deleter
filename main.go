package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"strings"
)

func main() {
	var token string

	// Arg parse
	flag.StringVar(&token, "t", "", "Bot token")
	flag.Parse()

	if strings.TrimSpace(token) == "" {
		fmt.Println("Missing bot token")
		os.Exit(1)
	}
	// Init discord client
	discord, err := discordgo.New("Bot " + "")

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("Bot is online")

	// Add all commands for bot
	var commandHandlers = []Handler{}

	// Register for events
	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) { OnMessageCreate(s, m, commandHandlers) })
}

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate, commands []Handler) {
	var args = strings.Split(strings.ToLower(m.Content), " ")

	if len(args) == 0 {
		return
	}

	// Iterate through command handlers till it finds one that is able to handle command
	for _, command := range commands {
		processedCommand, err := command.Handle(s, m, args)

		if err != nil {
			print(err)
		}

		if processedCommand {
			break
		}
	}

	// Built in commands
	// TODO: help, about
}
