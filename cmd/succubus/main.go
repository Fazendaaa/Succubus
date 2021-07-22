package main

import (
	"fmt"
	"os"
	"strings"

	succubus "github.com/Fazendaaa/Succubus/pkg"
	"github.com/spf13/cobra"
)

func main() {
	succCMD := make([](succubus.CMD), 2)
	appName := os.Getenv("SUCC_NAME")
	if "" == appName {
		appName = "succ"
	}
	rootCmd := &cobra.Command{
		Use: appName,
	}
	for _, command := range succCMD {
		translated := &cobra.Command{
			Use:   command.name,
			Short: command.usage.short,
			Long:  command.usage.long,
			Args:  cobra.MinimumNArgs(len(command.params)),
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Print: " + strings.Join(args, " "))
			},
		}

		rootCmd.AddCommand(translated)
	}
	rootCmd.Execute()
}
