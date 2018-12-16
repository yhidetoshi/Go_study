package main

import (
	"github.com/ashwanthkumar/slack-go-webhook"
	"os"
)

const (
	WEBHOOKURL = "https://hooks.slack.com/services/XXXXXXXXXXX"
	CHANNEL    = "dev"
	USERNAME   = "GoBot"
)

func main() {
	PostSlack("HelloWorld!!")
}

func PostSlack(msg string) {
	field1 := slack.Field{Title: "Message", Value: msg}

	attachment := slack.Attachment{}
	attachment.AddField(field1)
	color := "good"
	attachment.Color = &color
	payload := slack.Payload{
		Username:    USERNAME,
		Channel:     CHANNEL,
		Attachments: []slack.Attachment{attachment},
	}
	err := slack.Send(WEBHOOKURL, "", payload)
	if err != nil {
		os.Exit(1)
	}

}
