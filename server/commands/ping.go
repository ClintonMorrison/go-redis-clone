package commands

type PingCommand struct {}

func (cmd PingCommand) Apply(store Store) string {
	return "PONG"
}

func (cmd PingCommand) String() string {
	return "PING"
}
