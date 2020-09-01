package service

import (
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

// Server 起動
func Server() {
	logger := logrus.New()

	bot, err := linebot.New(secretToken, accessToken)
	if err != nil {
		logger.Fatalf("failed to set token to bot", err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/moneyforwardfree", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				logger.Errorf("request invalid signature: %v", err)
				w.WriteHeader(400)
			} else {
				logger.Errorf("failed to parse request: %v", err)
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					switch message.Text {
					case "image":
						logger.Infof("Received text image")
						GetAssetGraphImage()
						resp := linebot.NewImageMessage("https://drive.google.com/file/d/1EQ7_bZDq9TvYbYUX3-1mgIFdDzm9l-if/view?usp=sharing", "https://drive.google.com/file/d/1EQ7_bZDq9TvYbYUX3-1mgIFdDzm9l-if/view?usp=sharing")
						if _, err := bot.ReplyMessage(event.ReplyToken, resp).Do(); err != nil {
							logger.Errorf("failed to send image: %v", err)
						}
					default:
						logger.Infof("Received text")
						resp := linebot.NewTextMessage(message.Text)
						if _, err = bot.ReplyMessage(event.ReplyToken, resp).Do(); err != nil {
							logger.Errorf("failed to send text: %v", err)
						}
					}
				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		logger.Fatalf("failed to start server: %v", err)
	}
}
