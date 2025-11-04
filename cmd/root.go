package cmd

import (
	"fmt"

	"github.com/iamfittz/telegram-bot/bot"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "telegram-bot",
	Short: "Telegram bot powered by Go and Telebot",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting bot...")
		bot.Start()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
