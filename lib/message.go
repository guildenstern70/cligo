package lib

// Message Implements cli.Command and cli.Error interfaces
type Message struct {
	msg string
}

func (m *Message) Help() string {
	return m.msg
}

func (m *Message) Run(args []string) int {
	return 0
}

func (m *Message) Synopsis() string {
	return "Synopsis"
}

func (m *Message) Error() string {
	return "Message Error"
}

func NewMessage(message string) *Message {
	t := &Message{}
	t.msg = message
	return t
}
