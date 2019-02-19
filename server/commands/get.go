package commands

import "fmt"

type GetCommand struct {
	Key string
}

func (cmd GetCommand) Apply(store Store) string {
	return store.Get(cmd.Key).Serialize()
}

func (cmd GetCommand) String() string {
	return fmt.Sprintf("GET %s", cmd.Key)
}
