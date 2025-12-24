package bot

import (
	"context"
	"log"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	tele "gopkg.in/telebot.v4"
)

var tracer = otel.Tracer("telegram-bot")

func Start(ctx context.Context) {
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
		_, span := tracer.Start(ctx, "handle_message")
		defer span.End()

		user := c.Sender().FirstName
		msg := c.Text()

		span.SetAttributes(
			attribute.String("user.name", user),
			attribute.Int("message.length", len(msg)),
		)

		log.Printf("[TraceID: %s] Message from %s: %s",
			span.SpanContext().TraceID().String(), user, msg)

		return c.Send("Привіт, " + user + "! Ти написав: " + msg)
	})

	log.Println("Bot started!")
	b.Start()
}