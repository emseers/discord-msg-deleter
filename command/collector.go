package command

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

// MessageCollector
// /collect <message-snowflake-id> <condition-type> <condition-value>
// Condition type: number, time
// Condition value: # Messages to collect starting with ID or timestamp to collect till
type MessageCollector struct{}


func (h *MessageCollector) Name() string {
	return "Message Collector"
}

func (h *MessageCollector) Description() string {
	return "Collects messages starting from a given message ID to a given limit to local storage for further processing"
}

func (h *MessageCollector) Handle(s *discordgo.Session, m *discordgo.MessageCreate, args []string) (processedMessage bool, err error) {
	if args[0] != "/collect" {
		return
	}

	processedMessage = true

	if len(args) != 3 {
		err = errors.New("invalid arguments")
		return
	}

	numMessagesToCollect, err := strconv.ParseUint(args[2], 10, 32)

	if err != nil {
		err = fmt.Errorf("unable to parse value for number of messages to collect: %v", err)
		return
	}

	print(numMessagesToCollect)
	panic("implement me")
	return
}
