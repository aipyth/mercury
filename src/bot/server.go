package main

import (
	"log"
	"os"
	"time"
	// // "fmt"
	// "strconv"
	"net/http"
	tb "gopkg.in/tucnak/telebot.v2"
	"bot/dbutils"
)


func main() {
	serverMediaUrl := os.Getenv("SERVER_MEDIA_URL")
	localServerMediaUrl := "http://web:8000/media/"
	if os.Getenv("BOT_TOKEN") == "" {
		log.Println("No Token provided. Exiting...")
		return
	}

	// Connect to DB
	db := dbutils.Connect()
	defer db.Close()


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

	b.Handle("/roomtoday", func(m *tb.Message) {
		roomSlug := m.Payload
		room := dbutils.GetRoom(roomSlug)
		if room == nil {
			b.Send(m.Sender, "No such room")
			return
		}
		resp, err := http.Get(localServerMediaUrl + room.ScheduleImage)
		if err != nil {
			log.Println(err)
			return
		}
		photo := tb.Photo{
			File: tb.FromReader(resp.Body),
		}
		photo.Send(b, m.Sender, &tb.SendOptions{})
		
	})

	b.Handle("/today", func (m *tb.Message) {
		room := dbutils.GetRoom(m.Payload)
		if room == nil {
			b.Send(m.Sender, "No such room")
			return
		}
		b.Send(m.Sender, room.ScheduleToday().ToRepresentation(),
			&tb.SendOptions{
				ReplyTo: m,
				DisableWebPagePreview: true,
				ParseMode: tb.ModeHTML,
			})
	})

	b.Handle(tb.OnQuery, func (q* tb.Query) {
		log.Println("incoming query", *q)

		// results := make([]tb.Result, 2)
		// results[0] = &tb.ArticleResult{
		// 	Title: "Title",
		// 	Text: "Text",
		// }
		// results[1] = &tb.PhotoResult{
		// 	URL: "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/83/83418a5679456319c23f71f4431d3050d18eb42e_full.jpg",
		// 	ThumbURL: "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/83/83418a5679456319c23f71f4431d3050d18eb42e_full.jpg",
		// }
		
		// results[0].SetResultID(strconv.Itoa(1))
		// results[1].SetResultID(strconv.Itoa(0))
				
		// rooms := dbutils.GetRoom(q.Text)
		results := make([]tb.Result, 2)

		// for i, room := range rooms {
			// result := &tb.PhotoResult{
			// 	Title: room.Name,
			// 	// URL: "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/83/83418a5679456319c23f71f4431d3050d18eb42e_full.jpg",
			// 	// ThumbURL: "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/83/83418a5679456319c23f71f4431d3050d18eb42e_full.jpg",
			// 	URL: serverMediaUrl + room.ScheduleImage,
			// 	ThumbURL: serverMediaUrl + room.ScheduleImageThumb,
			// }
		// 	result := &tb.ArticleResult{
		// 		Title: room.Name,
		// 		URL: serverMediaUrl + room.ScheduleImage,
		// 		Text: room.Name,
		// 	}
		// 	result.SetResultID(strconv.Itoa(i))
		// 	log.Println("writing result", result)
		// 	results[i] = result
		// }
		results[0] = &tb.ArticleResult{
			Title: "Schedule for today",
		}

		log.Println(results)

		err := b.Answer(q, &tb.QueryResponse{
			Results: results,
			CacheTime: 1,
		})
		if err != nil {
			log.Println(err)	
		}
	})

	// b.Handle(tb.OnQuery, func (q* tb.Query) {

	// 	offset, err := strconv.ParseUint(q.Offset, 10, 64)
	// 	if err != nil {
	// 		offset = 0
	// 	}
	// 	var rooms []dbutils.Room
	// 	rooms = dbutils.GetRoomsBySlug(db, q.Text, offset)
	// 	log.Println("Got rooms ", rooms)
	// 	results := make([]tb.Result, len(rooms))
	// 	for i, v := range rooms {
	// 		if v.ScheduleImage == " " {
	// 			return
	// 		}
	// 		log.Println("writing room", v)
	// 		result := tb.PhotoResult{
	// 			Title: "Send " + v.Name + " schedule for today",
	// 			Description: v.Slug,
				
	// 			URL: serverMediaUrl + v.ScheduleImage,
	// 			ThumbURL: serverMediaUrl + v.ScheduleImageThumb,
	// 		}
	// 		log.Println("result is", result)
	// 		results[i] = &result
	// 		results[i].SetResultID(strconv.Itoa(i))
	// 	}
	// 	resp := tb.QueryResponse{
	// 		Results:   	results,
	// 		CacheTime: 	1,
	// 		NextOffset: strconv.FormatUint(offset+1, 10),
	// 	}
	// 	log.Println(resp)
	// 	err = b.Answer(q, &resp)
	// 	if err != nil {
	// 		log.Println(err)	
	// 	}
	// })

	b.Start()

}
