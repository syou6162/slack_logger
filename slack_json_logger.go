package main

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/nlopes/slack"
)

type Message struct {
	Username string `json:"username"`
	Date     string `json:"date"`
	Channel  string `json:"channel"`
	Team     string `json:"team"`
	Text     string `json:"text"`
	Ts       string `json:"ts"`
}

func getUserId2UserNameMap(users []slack.User) map[string]string {
	m := make(map[string]string)
	for _, u := range users {
		m[u.ID] = u.Name
	}
	return m
}

const layout = "20060102T15:04:05-07:00"

func main() {
	api := slack.New(os.Args[1])
	teamInfo, _ := api.GetTeamInfo()
	users, _ := api.GetUsers()
	id2name := getUserId2UserNameMap(users)

	channels, _ := api.GetChannels(true)

	for _, c := range channels {
		if c.IsMember {
			history, _ := api.GetChannelHistory(c.ID, slack.NewHistoryParameters())
			for _, m := range history.Messages {
				u, _ := strconv.ParseFloat(m.Timestamp, 64)
				t := time.Unix(int64(u), 0)
				mess := Message{}
				mess.Username = id2name[m.User]
				mess.Date = t.Format(layout)
				mess.Channel = c.Name
				mess.Team = teamInfo.Domain
				mess.Ts = m.Timestamp
				mess.Text = m.Text
				json.NewEncoder(os.Stdout).Encode(mess)
			}
		}
	}
}
