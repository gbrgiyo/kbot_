/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	TeleToken string = os.Getenv("TELE_TOKEN")
)

var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "KBot is a bot for Telegram",
	Long: `KBot is a fully functional bot for Telegram.

Created for fun and learning purposes. The main goal is to learn Go and Telegram Bot API.
It's a work in progress, so expect some bugs and missing features.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kbot %s started\n", appVersion)
		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Please check your TELE_TOKEN env variable, %s", err)
		}

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Printf(m.Message().Payload, m.Text())
			payload := m.Message().Payload

			switch payload {
			case "hello":
				err = m.Send((fmt.Sprintf("Hello! I'm kbot %s", appVersion)))
			case "help":
				err = m.Send((fmt.Sprintf("I'm kbot %s, I can't help you yet", appVersion)))
			case "ping":
				err = m.Send(("pong"))
			default:
				err = m.Send(("Sorry, I don't understand"))

			}

			return err
		})
		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)

}
