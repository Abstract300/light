package messages

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

type MessageState struct {
	Store Database
}

type Message struct {
	ID              string
	ChannelID       string
	GuildID         string
	Timestamp       string
	EditedTimestamp string
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
		ID:            m.ID,
		ChannelID:     m.ChannelID,
		GuildID:       m.GuildID,
		MessageAuthor: user,
	}

	discordMsgTime, err := m.Timestamp.Parse()
	if err != nil {
		log.Println(err)
	}
	userMsg.Timestamp = discordMsgTime.Format(time.RFC1123Z)

	msg.Store.StoreChannelMessage(userMsg)
}
