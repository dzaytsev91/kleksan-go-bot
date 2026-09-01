package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
	}

	b, err := bot.New(os.Getenv("BOT_TOKEN"), opts...)
	if nil != err {
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/where", bot.MatchTypeExact, whereHandler)

	b.Start(ctx)
}

func whereHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	// Счёт дней с Unix-эпохи растёт на 1 каждый день независимо от границ
	// месяца и года, поэтому соседние дни всегда имеют разную чётность и
	// направление никогда не повторяется (в отличие от Day()/YearDay()).
	dayNumber := int(time.Now().Unix() / 86400)

	if dayNumber%2 != 0 {
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      "Сегодня колоть влево",
			ParseMode: models.ParseModeMarkdown,
			ReplyParameters: &models.ReplyParameters{
				MessageID: update.Message.ID,
				ChatID:    update.Message.Chat.ID,
			},
		})
		if err != nil {
			panic(err)
		}
	} else {
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      "Сегодня колоть вправо",
			ParseMode: models.ParseModeMarkdown,
			ReplyParameters: &models.ReplyParameters{
				MessageID: update.Message.ID,
				ChatID:    update.Message.Chat.ID,
			},
		})
		if err != nil {
			panic(err)
		}
	}
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Чтобы узнать куда колоть сегодня нажми команду /where",
	})
	if err != nil {
		panic(err)
	}
}
