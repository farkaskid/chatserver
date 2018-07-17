package message

// Message represents a message
type Message struct {
	Sender, Content string
	Recipients      []string
}
