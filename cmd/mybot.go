/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var(
	Teletoken = os.Getenv("TELE_TOKEN")
)

// mybotCmd represents the mybot command
var mybotCmd = &cobra.Command{
	Use:   "mybot",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mybot called")
	},
}

func init() {
	rootCmd.AddCommand(mybotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mybotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mybotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
