package handler

import (
	"encoding/json"
	"strconv"

	tele "gopkg.in/telebot.v3"
)

var storedPoll *tele.Message

func (h Handler) CreatePoll(c tele.Context) error {
	newPoll := &tele.Poll{
		ID:       strconv.Itoa(1),
		Question: "TEST",
		Options: []tele.PollOption{
			{Text: "1"},
			{Text: "2"},
			{Text: "3"},
		},
	}

	poll, err := c.Bot().Send(c.Sender(), newPoll)
	storedPoll = poll
	return err
}

func (h Handler) ClosePoll(c tele.Context) error {
	poll, err := c.Bot().StopPoll(storedPoll)
	if err != nil {
		return err
	}

	result, err := json.Marshal(poll)
	if err != nil {
		return err
	}
	return c.Send(string(result))
}
