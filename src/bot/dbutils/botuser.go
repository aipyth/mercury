package dbutils

import (
	"log"
	"strings"
	"database/sql"
)

type BotUser struct {
	ChatId int32
	rooms string
	UserId sql.NullInt64
	DefaultRoom sql.NullString

	Rooms []string
}

func CreateTableIfNotExistsBotUsers() {
	// global DB must be set
	if DB == nil {
		log.Println("Global DB is not set to create table bot_user")
		return
	}
	DB.Exec(`CREATE TABLE IF NOT EXISTS bot_user (
		user_id integer,
		chat_id integer UNIQUE NOT NULL,
		rooms text DEFAULT '',
		default_room text,
		FOREIGN KEY (user_id)
			REFERENCES users_customuser(id)
	);`)
}

func GetBotUser(chatId int) *BotUser {
	botuser := &BotUser{}
	err := DB.QueryRow(`SELECT chat_id, rooms, user_id, default_room
						FROM bot_user
						WHERE chat_id=$1
						;`, chatId).Scan(
						&botuser.ChatId, &botuser.rooms, &botuser.UserId,
						&botuser.DefaultRoom)
	// if err == sql.ErrNoRows {
	// 	return nil
	// }
	if err != nil {
		log.Println(err)
		return nil
	}
	botuser.Rooms = strings.Split(botuser.rooms, ";")
	if botuser.Rooms[0] == "" {
		botuser.Rooms = make([]string, 0)
	}
	return botuser
}

func CreateBotUser(chatId int) *BotUser {
	_, err := DB.Exec(`INSERT INTO bot_user(chat_id)
			 VALUES ($1)`, chatId)
	if err != nil {
		log.Println(err)
	}
	return GetBotUser(chatId)
}

func CreateBotUserIfNotExists(chatId int) *BotUser {
	var count int64
	DB.QueryRow(`SELECT COUNT(*)
				 FROM bot_user
				 WHERE chat_id=$1
				 ;`,chatId).Scan(&count)
	if count == 0 {
		return CreateBotUser(chatId)
	}
	return GetBotUser(chatId)
}

func (u *BotUser) AddUsedRoom(room *Room) {
	for i, v := range u.Rooms {
		if v == "" {
			// delete empty string from list if present
			u.Rooms[i] = u.Rooms[len(u.Rooms)-1]
			u.Rooms[len(u.Rooms)-1] = ""
			u.Rooms = u.Rooms[:len(u.Rooms)-1]
		}
		if v == room.Slug {
			// we delete this record to ...
			u.Rooms[i] = u.Rooms[len(u.Rooms)-1]
			u.Rooms[len(u.Rooms)-1] = ""
			u.Rooms = u.Rooms[:len(u.Rooms)-1]
			// return
		}
	}
	// ... to append it to the beginning again
	// this way we realize last searches system
	u.Rooms = append(u.Rooms, "")
    copy(u.Rooms[1:], u.Rooms)
    u.Rooms[0] = room.Slug
	u.rooms = strings.Join(u.Rooms, ";")
	_, err := DB.Exec(`UPDATE bot_user
					   SET rooms=$2
					   WHERE chat_id=$1
					   ;`, u.ChatId, u.rooms)
	if err != nil {
		log.Println(err)
	}
}

func (u *BotUser) SetDefaultRoom(roomSlug string) bool {
	_, err := DB.Exec(`UPDATE bot_user
					   SET default_room=$1
					   WHERE chat_id=$2
					   ;`, roomSlug, u.ChatId)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}