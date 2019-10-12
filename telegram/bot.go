package telegram

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mcctor/redditart/db"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	Bot *tb.Bot
)

const (
	redditBaseUrl = "https://reddit.com"
)

func init() {
	var err error
	Bot, err = tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
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
		_, err := Bot.Send(message.Sender, "Welcome back! Press /startstreaming to start receiving images")
		checkError(err)

	} else {
		db.AddUser(db.User{
			UserID:    message.Sender.ID,
			FirstName: message.Sender.FirstName,
		})
		_, err := Bot.Send(message.Sender, "Welcome! Press /startstreaming to start receiving images")
		checkError(err)
	}
}

func startStopStreamingHelper(message *tb.Message, streaming bool, streamingText, notStreamingText string) {
	user, exists := db.GetUserByID(message.Sender.ID)
	registerAccountFirst(exists, message.Sender)
	if user.Streaming {
		_, err := Bot.Send(message.Sender, streamingText)
		checkError(err)
	} else {
		user.Streaming = streaming
		_, err := Bot.Send(
			message.Sender,
			notStreamingText)
		checkError(err)
		db.UpdateUser(user)
	}
}

func startstreamingHandler(message *tb.Message) {
	startStopStreamingHelper(
		message,
		true,
		"You are already streaming",
		"You will now periodically receive new images, to stop, press /stopstreaming.")
}

func stopstreamingHandler(message *tb.Message) {
	startStopStreamingHelper(
		message,
		true,
		"You will no longer receive any new images. Press /startstreaming to get images.",
		"You are currently not streaming.")
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

func registerAccountFirst(userExists bool, recipient tb.Recipient) {
	if !userExists {
		_, err := Bot.Send(recipient, "Register account first by pressing /start")
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}
