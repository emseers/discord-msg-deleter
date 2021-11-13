package main

import "github.com/bwmarrin/discordgo"

type Handler interface {
	Name() string
	Description() string
	Handle(s *discordgo.Session, m *discordgo.MessageCreate, args []string) (bool, error)
}
