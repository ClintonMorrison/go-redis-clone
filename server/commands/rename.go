package commands

import (
	"fmt"
)

type RenameCommand struct {
	SrcKey string
	DestKey string
}

func (cmd RenameCommand) Apply(store Store) string {
	value := store.Get(cmd.SrcKey)
	store.Set(cmd.DestKey, value) // TODO: copy expiry date?
	store.Delete(cmd.SrcKey)
	return ""
}

func (cmd RenameCommand) String() string {
	return fmt.Sprintf("SET %s %s", cmd.SrcKey, cmd.DestKey)
}
