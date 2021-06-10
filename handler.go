package main

import "github.com/bwmarrin/discordgo"

type Handler interface {
	Name() string
	Description() string
	Handle(s *discordgo.Session, m *discordgo.MessageCreate, cmdSlices []string) (bool, error)
}
