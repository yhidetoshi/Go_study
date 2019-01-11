package main

import (
	"github.com/ashwanthkumar/slack-go-webhook"
	"os"
)

const (
	WEBHOOKURL = "https://hooks.slack.com/services/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"
	CHANNEL    = "dev"
	USERNAME   = "GoBot"
)

func main() {
	PostSlack("HelloWorld!!")
}

func PostSlack(msg string) {
	field1 := slack.Field{Title: "Message", Value: msg}
	field2 := slack.Field{Title: "AnythingKey", Value: "AnythingValue"}

	attachment := slack.Attachment{}
	attachment.AddField(field1).AddField(field2)
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
