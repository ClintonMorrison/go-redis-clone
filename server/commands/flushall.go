package commands

type FlushAllCommand struct {}

func (cmd FlushAllCommand) Apply(store Store) string {
	store.DeleteAll()
	return ""
}

func (cmd FlushAllCommand) String() string {
	return "KEYS"
}
