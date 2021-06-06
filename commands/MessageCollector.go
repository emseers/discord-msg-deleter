package commands

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

// MessageCollector
// /collect <message-snowflake-id> <condition-type> <condition-value>
// Condition type: number, time
// Condition value: # Messages to collect starting with ID or timestamp to collect till
type MessageCollector struct{}


func (h MessageCollector) Name() string {
	return "Message Collector"
}

func (h MessageCollector) Description() string {
	return "Collects messages starting from a given message ID to a given limit to local storage for further processing"
}

func (h MessageCollector) Execute(s *discordgo.Session, m *discordgo.MessageCreate, cmdSlices []string) (bool, error) {
	if cmdSlices[0] != "collect" {
		return false, nil
	}

	if len(cmdSlices) != 3 {
		return true, errors.New("invalid arguments")
	}

	numMessagesToCollect, err := strconv.ParseUint(cmdSlices[2], 10, 32)

	if err != nil {
		return true, errors.New("invalid value for number of messages to collect")
	}

	print(numMessagesToCollect)
	panic("implement me")
}
