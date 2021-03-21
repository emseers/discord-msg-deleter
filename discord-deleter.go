package main

import (
	"discord-msg-deleter/commands"
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
	var commandHandlers = []commands.ICommandHandler{}

	// Register for events
	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) { OnMessageCreate(s, m, commandHandlers) })
}

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate, commands []commands.ICommandHandler) {
	var cmdSlices = strings.Split(strings.ToLower(m.Content), "")

	for _, command := range commands {
		canExecute, err := command.Execute(s, m, cmdSlices)

		if err != nil {
			print(err)
		}

		if canExecute {
			break
		}
	}

	// Built in commands
	// TODO: help, about
}
