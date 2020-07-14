package store

import (
	"log"

	"github.com/abstract300/light/messages"
)

func (st *Store) StoreChannelMessage(msg messages.Message) error {
	_, err := st.DB.Query("SELECT id FROM messageauthor WHERE id = $1", msg.MessageAuthor.ID)
	if err != nil {
		log.Println("[Author ID not in the DB]", err)

		_, err = st.DB.Exec("INSERT INTO messageauthor(id, username) VALUES($1, $2)",
			msg.MessageAuthor.ID, msg.MessageAuthor.Username)

		if err != nil {
			log.Println("[Author create error]", err)
		}
	}

	_, err = st.DB.Exec(`INSERT INTO usermessage(authorid, channelid, guildid, timestamp,  messageid)
			VALUES($1, $2, $3, $4, $5)`,
		msg.MessageAuthor.ID, msg.ChannelID, msg.GuildID, msg.Timestamp, msg.ID)

	if err != nil {
		log.Println("[Message insert error]", err)
	}

	return err
}
