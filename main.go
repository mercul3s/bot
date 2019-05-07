package main

import (
	"os"
	"slackbot/db"
	"slackbot/users"
	"strings"

	"github.com/sbstjn/hanu"
	"github.com/sirupsen/logrus"
)

type config struct {
	Logger *logrus.Logger
}

func new() *config {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	return &config{
		Logger: log,
	}
}

func main() {
	Version := "0.0.1"

	c := new()
	bot, err := hanu.New(os.Getenv("SLACK_TOKEN"))
	if err != nil {
		c.Logger.Fatal(err)
	}

	store, err := db.Connect("bot.db")
	if err != nil {
		c.Logger.Fatal(err)
	}

	u := users.New(store)

	bot.Command("add <user> as type <type>", func(conv hanu.ConversationInterface) {
		user, _ := conv.String("user")
		err := u.Add(user)
		c.Logger.Info("added user")
	})

	bot.Command("shout <word>", func(conv hanu.ConversationInterface) {
		word, _ := conv.String("word")
		c.Logger.Info("replied to user")
		conv.Reply(strings.ToUpper(word))
	})

	bot.Command("whisper <word>", func(conv hanu.ConversationInterface) {
		word, _ := conv.String("word")
		conv.Reply(strings.ToLower(word))
	})

	bot.Command("version", func(conv hanu.ConversationInterface) {
		conv.Reply("Thanks for asking! I'm running `%s`", Version)
	})

	bot.Listen()
}
