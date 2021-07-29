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
	rootCmd := &cobra.Command{
		Use: appName,
	}

	if "" == appName {
		appName = "succ"
	}

	for _, command := range succCMD {
		translated := &cobra.Command{
			Use:   command.Name,
			Short: command.Usage.Short,
			Long:  command.Usage.Long,
			Args:  cobra.MinimumNArgs(command.Args),
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Print: " + strings.Join(args, " "))
				command.Function(args)
			},
		}

		rootCmd.AddCommand(translated)
	}

	rootCmd.Execute()
}
