package commands

import "github.com/bwmarrin/discordgo"

type ICommandHandler interface {
	Name() string
	Description() string
	Execute(s *discordgo.Session, m *discordgo.MessageCreate, cmdSlices []string) (bool, error)
}
