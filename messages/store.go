package messages

type Database interface {
	StoreChannelMessage(msg Message) error
}
