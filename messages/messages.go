package messages

import (
	"github.com/bwmarrin/discordgo"
)

type MessageState struct {
	Store Database
}

type Message struct {
	ID              string
	ChannelID       string
	GuildID         string
	Timestamp       discordgo.Timestamp
	EditedTimestamp discordgo.Timestamp
	MessageAuthor   Author
}

type Author struct {
	ID       string
	Username string
}

func NewMessageState(db Database) *MessageState {
	return &MessageState{
		Store: db,
	}
}

func (msg MessageState) MessageCreateEvent(s *discordgo.Session, m *discordgo.MessageCreate) {

	user := Author{
		ID:       m.Author.ID,
		Username: m.Author.Username,
	}
	userMsg := Message{
		ID:              m.ID,
		ChannelID:       m.ChannelID,
		GuildID:         m.GuildID,
		Timestamp:       m.Timestamp,
		EditedTimestamp: m.EditedTimestamp,
		MessageAuthor:   user,
	}

	msg.Store.StoreChannelMessage(userMsg)
}
