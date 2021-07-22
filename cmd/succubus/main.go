package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Fazendaaa/Succubus/pkg"
	"github.com/spf13/cobra"
)

func cmdToCobra(cmd CMD) (ok bool, fail error) {
	return ok, fail
}

func main() {
	cmdPrint := &cobra.Command{
		Use:   "print [string to print]",
		Short: "Print anything to the screen",
		Long: `print is for printing anything back to the screen.
For many years people have printed back to the screen.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	cmdEcho := &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Long: `echo is for echoing anything back.
Echo works a lot like print, except it has a child command.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}

	appName := os.Getenv("SUCC_NAME")
	if "" == appName {
		appName = "succ"
	}
	rootCmd := &cobra.Command{
		Use: appName,
	}

	rootCmd.AddCommand(cmdPrint, cmdEcho)
	rootCmd.Execute()
}
