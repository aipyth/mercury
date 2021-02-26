package main

import (
	"log"
	"os"
	"time"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	if os.Getenv("BOT_TOKEN") == "" {
		return
	}

	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hello World!")
	})

	b.Handle(tb.OnQuery, func (q* tb.Query) {
		if q.Text != "" {
			log.Println("Getting concrete room")
		} else {
			// empty request so we list awailable rooms
			log.Println("List awailable rooms")
		}
	})

	b.Start()

}
