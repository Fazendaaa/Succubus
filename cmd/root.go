package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "succubus",
	Short: "Succubus is project manager",
	Long: `Succubus is made with Go to help you handle all of your project needs.
Complete documentation is available at https://github.com/Fazendaaa/Shojo

Succubus is also part of the Container tooling For Developers (CFD) initiative:
https://github.com/Fazendaaa/CFD`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
