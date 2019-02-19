package commands

import (
	"go-redis-clone/server/primitives"
	"fmt"
)

type SetCommand struct {
	Key string
	Value primitives.Value
}

func (cmd SetCommand) Apply(store Store) string {
	store.Set(cmd.Key, cmd.Value)
	return ""
}

func (cmd SetCommand) String() string {
	return fmt.Sprintf("SET %s %s", cmd.Key, cmd.Value.Serialize())
}
