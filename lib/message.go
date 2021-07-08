package lib

// Message Implements cli.Command and cli.Error interfaces
type Message struct {
	msg string
}

func (m *Message) Help() string {
	return m.msg
}

func (m *Message) Run(args []string) int {
	println("cligo message")
	println("")
	println("Available commands are:")
	println("\tdefault\t\t\tGet the default message")
	println("\tspecial --id [id]\tGet the special message for the ID=id")
	return 0
}

func (m *Message) Synopsis() string {
	return "Get a default or a special message"
}

func NewMessage(message string) *Message {
	t := &Message{}
	t.msg = message
	return t
}
