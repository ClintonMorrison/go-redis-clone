package commands

import "fmt"

type DeleteCommand struct {
	Key string
}

func (cmd DeleteCommand) Apply(store Store) string {
	return store.Delete(cmd.Key).Serialize()
}

func (cmd DeleteCommand) String() string {
	return fmt.Sprintf("DELETE %s", cmd.Key)
}