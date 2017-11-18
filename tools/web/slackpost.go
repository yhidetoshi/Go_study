package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"fmt"
)

var (
	WebhookUrl string = "https://hooks.slack.com/services/XXXXX"
)

type Slack struct {
	Text        string `json:"text"`
	Username    string `json:"username"`
	Icon_emoji  string `json:"icon_emoji"`
	Icon_url    string `json:"icon_url"`
	Channel     string `json:"channel"`
}

func main() {
	arg := os.Args[1:]
	params, _ := json.Marshal(Slack{
		fmt.Sprintf("%s", arg),
		"Bot",
		"",
		"http://ww.techscore.com/blog/wp/wp-content/uploads/2016/12/gopher_ueda.png",
		"channelname"})

	res, _ := http.PostForm(
		WebhookUrl,
		url.Values{"payload": {string(params)}},
	)
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	println(string(body))
}
