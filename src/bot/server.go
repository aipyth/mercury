package main

import (
	"log"
	"os"
	"time"
	// // "fmt"
	"strconv"
	"net/http"
	tb "gopkg.in/tucnak/telebot.v2"
	"bot/dbutils"
)


func main() {
	serverMediaUrl := os.Getenv("SERVER_MEDIA_URL")
	localServerMediaUrl := "http://web:8000/media/"
	if os.Getenv("BOT_TOKEN") == "" {
		log.Println("No Token provided. Exiting...")
		os.Exit(1)
	}

	// Connect to DB
	db := dbutils.Connect()
	defer db.Close()
	dbutils.CreateTableIfNotExistsBotUsers()


	// Start bot
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	log.Println("GoBot started")
	log.Println("servermeidaurl ", serverMediaUrl)

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/gachi", func(m *tb.Message) {
		photo := tb.Photo{
			File: tb.FromURL("https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/83/83418a5679456319c23f71f4431d3050d18eb42e_full.jpg"),
		}
		if m.Sender != nil {
			photo.Send(b, m.Sender, &tb.SendOptions{})
		} else {
			photo.Send(b, m.Chat, &tb.SendOptions{})
		}
	})

	b.Handle("/todayim", func(m *tb.Message) {
		buser := dbutils.CreateBotUserIfNotExists(m.Sender.ID)
		var to tb.Recipient
		if m.Sender == nil {
			to = m.Chat	
		} else {
			to = m.Sender
		}
		room := dbutils.GetRoom(m.Payload)
		if room == nil {
			if buser.DefaultRoom.Valid {
				room = dbutils.GetRoom(buser.DefaultRoom.String)
			} else {
				b.Send(to, "you have not set a default room.\ntry now with /bind\nor you can specify room code after /todayim")
				return
			}
		}
		resp, err := http.Get(localServerMediaUrl + room.ScheduleImage)
		if err != nil {
			log.Println(err)
			b.Send(to, "rain on clouds! cannot get image..")
		} else {
			photo := tb.Photo{
				File: tb.FromReader(resp.Body),
			}
			photo.Send(b, to, &tb.SendOptions{})
		}
	})

	// b.Handle("/tomorrowim", func(m *tb.Message) {
	// 	buser := dbutils.CreateBotUserIfNotExists(m.Sender.ID)
	// 	var to tb.Recipient
	// 	if m.Sender == nil {
	// 		to = m.Chat	
	// 	} else {
	// 		to = m.Sender
	// 	}
	// 	room := dbutils.GetRoom(m.Payload)
	// 	if room == nil {
	// 		if buser.DefaultRoom.Valid {
	// 			room = dbutils.GetRoom(buser.DefaultRoom.String)
	// 		} else {
	// 			b.Send(to, "you have not set a default room.\ntry now with /bind\nor you can specify room code after /tomorrowim")
	// 			return
	// 		}
	// 	}
	// 	resp, err := http.Get(localServerMediaUrl + room.ScheduleImage)
	// 	if err != nil {
	// 		log.Println(err)
	// 		b.Send(to, "rain on clouds! cannot get image..")
	// 	} else {
	// 		photo := tb.Photo{
	// 			File: tb.FromReader(resp.Body),
	// 		}
	// 		photo.Send(b, to, &tb.SendOptions{})
	// 	}
	// })

	b.Handle("/today", func (m *tb.Message) {
		buser := dbutils.CreateBotUserIfNotExists(m.Sender.ID)
		var to tb.Recipient
		if m.Sender == nil {
			to = m.Chat	
		} else {
			to = m.Sender
		}
		room := dbutils.GetRoom(m.Payload)
		if room == nil {
			if buser.DefaultRoom.Valid {
				room := dbutils.GetRoom(buser.DefaultRoom.String)
				b.Send(to, room.ScheduleToday().ToRepresentation(),
				&tb.SendOptions{
					ReplyTo: m,
					DisableWebPagePreview: true,
					ParseMode: tb.ModeHTML,
				})
			} else {
				b.Send(to, "you have not set a default room.\ntry now with /bind\nor you can specify room code after /today")
			}
		} else {
			buser.AddUsedRoom(room)
			b.Send(to, room.ScheduleToday().ToRepresentation(),
				&tb.SendOptions{
					ReplyTo: m,
					DisableWebPagePreview: true,
					ParseMode: tb.ModeHTML,
				})
		}
	})

	b.Handle("/bind", func (m *tb.Message) {
		buser := dbutils.CreateBotUserIfNotExists(m.Sender.ID)
		var to tb.Recipient
		if m.Sender == nil {
			// we're in chat
			to = m.Chat
		} else {
			to = m.Sender
		}
		if m.Payload == "" {
			b.Send(to, "write room code after /bind")
			return
		}
		room := dbutils.GetRoom(m.Payload)
		if room == nil {
			b.Send(to, "No such room " + m.Payload)
			// return
		} else {
			if buser.SetDefaultRoom(m.Payload) {
				b.Send(to, "okk, set!")
			} else {
				b.Send(to, "something wrong. cannot set now..")
			}
		}
	})

	b.Handle("/tomorrow", func (m *tb.Message) {
		buser := dbutils.CreateBotUserIfNotExists(m.Sender.ID)
		
		var to tb.Recipient
		if m.Sender == nil {
			to = m.Chat	
		} else {
			to = m.Sender
		}
		room := dbutils.GetRoom(m.Payload)

		if room == nil {
			if buser.DefaultRoom.Valid {
				room := dbutils.GetRoom(buser.DefaultRoom.String)
				b.Send(to, room.ScheduleTomorrow().ToRepresentation(),
				&tb.SendOptions{
					ReplyTo: m,
					DisableWebPagePreview: true,
					ParseMode: tb.ModeHTML,
				})
			} else {
				b.Send(to, "you have not set a default room.\ntry now with /bind\nor you can specify room code after /tomorrow")
			}
		} else {
			buser.AddUsedRoom(room)
			b.Send(to, room.ScheduleTomorrow().ToRepresentation(),
				&tb.SendOptions{
					ReplyTo: m,
					DisableWebPagePreview: true,
					ParseMode: tb.ModeHTML,
				})
		}
	})

	b.Handle(tb.OnQuery, func (q* tb.Query) {
		buser := dbutils.CreateBotUserIfNotExists(q.From.ID)

		log.Println("incoming query", *q)
		var results []tb.Result
		if q.Text == "" {
			results = make([]tb.Result, 0)
			for i, roomSlug := range buser.Rooms {
				room := dbutils.GetRoom(roomSlug)
				results = append(results, &tb.ArticleResult{
					Title: "Schedule " + room.Name,
				})
				results[i].SetContent(&tb.InputTextMessageContent{
					Text: room.ScheduleToday().ToRepresentation(),
					ParseMode: tb.ModeHTML,
				})
				results[i].SetResultID(strconv.Itoa(i))
			}
		} else {

			room := dbutils.GetRoom(q.Text)
			if room == nil {
				results = make([]tb.Result, 1)
				results[0] = &tb.ArticleResult{
					Title: "No such room",
					Text: "No such room " + q.Text,
				}
				results[0].SetResultID(strconv.Itoa(1))
			} else {
				buser.AddUsedRoom(room)
				results = make([]tb.Result, 2)
				results[0] = &tb.ArticleResult{
					Title: "Schedule for today",
				}
				results[0].SetContent(&tb.InputTextMessageContent{
					Text: room.ScheduleToday().ToRepresentation(),
					ParseMode: tb.ModeHTML,
				})
				results[1] = &tb.PhotoResult{
					Title: "Schedule image today",
					URL: serverMediaUrl + room.ScheduleImage,
					ThumbURL: serverMediaUrl + room.ScheduleImage,
				}
				results[0].SetResultID(strconv.Itoa(1))
				results[1].SetResultID(strconv.Itoa(2))	
			}
		}
		
		err := b.Answer(q, &tb.QueryResponse{
			Results: results,
			CacheTime: 1,
		})
		if err != nil {
			log.Println(err)	
		}
	})


	b.Start()

}
