package gohelpers

import goteamsnotify "github.com/atc0005/go-teams-notify"

func TeamsSuccess(message string, title string, channel string) {
	mstClient := goteamsnotify.NewClient()
	msgCard := goteamsnotify.NewMessageCard()
	msgCard.Title = title
	msgCard.Text = message
	msgCard.ThemeColor = "#00FF04"
	mstClient.Send(channel, msgCard)
}

func TeamsError(err error, title string, channel string) {
	mstClient := goteamsnotify.NewClient()
	msgCard := goteamsnotify.NewMessageCard()
	msgCard.Title = title
	msgCard.Text = "Error message: " + err.Error()
	msgCard.ThemeColor = "#FF0000"
	mstClient.Send(channel, msgCard)
}

func TeamsInfo(message string, title string, channel string) {
	mstClient := goteamsnotify.NewClient()
	msgCard := goteamsnotify.NewMessageCard()
	msgCard.Title = title
	msgCard.Text = message
	msgCard.ThemeColor = "#ffe600"
	mstClient.Send(channel, msgCard)
}
