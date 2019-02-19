package commands

import (
	"fmt"
	"go-redis-clone/server/primitives"
	"strings"
)

type ParseError struct {
	msg string
}

func (p *ParseError) Error() string {
	return p.msg
}

func verifyArgumentCount(expectedCount int, tokens []string) *ParseError {
	if len(tokens) != (expectedCount + 1) {
		msg := fmt.Sprintf("expected %d arguments", expectedCount)
		return 	&ParseError{msg}
	}

	return nil
}

var argumentCounts = map[string]int{
	"PING": 0,
	"SET":   2,
	"GET": 1,
	"DELETE": 1,
	"KEYS": 0,
	"FLUSHALL": 0,
	"RENAME": 2,
}

func ParseCommand(message string) (Command, error) {
	tokens := Tokenize(message)
	// Make sure message not empty
	if len(tokens) == 0 {
		return nil, &ParseError{"No tokens in message"}
	}

	// Make sure command exists
	command := strings.ToUpper(tokens[0])
	requiredCount, cmdFound := argumentCounts[command]
	if !cmdFound {
		return nil, &ParseError{"Unrecognized command"}
	}

	// Make sure correct number of arguments given
	err := verifyArgumentCount(requiredCount, tokens)
	if err != nil {
		return nil, err
	}

	// Parse into command struct
	switch command {
	case "PING":
		return PingCommand{}, nil

	case "SET":
		return SetCommand{
			tokens[1],
			primitives.Value{ tokens[2], "string"},
		}, nil

	case "GET":
		return GetCommand{
			tokens[1],
		}, nil

	case "DELETE":
		return DeleteCommand{
			tokens[1],
		}, nil

	case "RENAME":
		return RenameCommand{
			tokens[1],
			tokens[2],
		}, nil

	case "KEYS":
		return KeysCommand{}, nil

	case "FLUSHALL":
		return FlushAllCommand{}, nil
	}

	return nil, nil
}
