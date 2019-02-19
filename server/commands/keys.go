package commands

import (
	"encoding/json"
)

type KeysCommand struct {}

func (cmd KeysCommand) Apply(store Store) string {
	bytes, _ := json.Marshal(store.Keys())
	return string(bytes)
}

func (cmd KeysCommand) String() string {
	return "KEYS"
}
