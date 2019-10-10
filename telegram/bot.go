package telegram

import (
	"fmt"
	"log"
	"time"

	"github.com/mcctor/redditart/db"
	tb "gopkg.in/tucnak/telebot.v2"
)

var Bot *tb.Bot

const (
	redditBaseUrl = "https://reddit.com"
)

func init() {
	var err error
	Bot, err = tb.NewBot(tb.Settings{
		Token:  "743129363:AAGhApUDIdj4Khk9CibgRTCHWt0BMojofMo",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	checkError(err)

	// register bot handlers here
	Bot.Handle("/start", startHandler)
	Bot.Handle("/startstreaming", startstreamingHandler)
	Bot.Handle("/stopstreaming", stopstreamingHandler)
}

func startHandler(message *tb.Message) {
	if _, exists := db.GetUserByID(message.Sender.ID); exists {
		// welcome old user
		_, err := Bot.Send(message.Sender, "Welcome back! Press /startstreaming to start receiving images")
		checkError(err)

	} else {
		// register new user
		db.AddUser(db.User{
			UserID:    message.Sender.ID,
			FirstName: message.Sender.FirstName,
		})

		// welcome new user
		_, err := Bot.Send(message.Sender, "Welcome! Press /startstreaming to start receiving images")
		checkError(err)
	}
}

func startstreamingHandler(message *tb.Message) {
	user, _ := db.GetUserByID(message.Sender.ID)
	user.Streaming = true
	db.UpdateUser(user)
}

func stopstreamingHandler(message *tb.Message) {
	user, _ := db.GetUserByID(message.Sender.ID)
	user.Streaming = false
	db.UpdateUser(user)
}

func sendPhoto(user *tb.User, newPost db.Post) {
	postUrl := redditBaseUrl + newPost.Link
	caption := fmt.Sprintf("\n🏞 _%s_\n👤 by #%s\n🌏 [Reddit](%s)\n", newPost.Caption, newPost.Author, postUrl)
	photo := &tb.Photo{File: tb.FromURL(newPost.ImageUrl), Caption: caption}
	_, err := Bot.Send(user, photo, tb.ModeMarkdown)
	checkError(err)
}

func SendPhotoToAll(newPost db.Post) {
	for _, user := range db.GetAllUsers(true) {
		sendPhoto(&tb.User{ID: user.UserID}, newPost)
	}
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}
