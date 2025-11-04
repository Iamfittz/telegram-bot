package bot

import (
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

func Start() {
	token := os.Getenv("TELE_TOKEN")
	if token == "" {
		log.Fatal("TELE_TOKEN not set")
	}

	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	b.Handle(tele.OnText, func(c tele.Context) error {
		user := c.Sender().FirstName
		msg := c.Text()
		log.Printf("Message from %s: %s", user, msg)
		return c.Send("Привіт, " + user + "! Ти написав: " + msg)
	})

	log.Println("Bot started!")
	b.Start()
}
