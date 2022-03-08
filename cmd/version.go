package cmd

import (
	"fmt"

	"github.com/Fazendaaa/Succubus/controllers"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Removes a packages/libraries from the current project",
	Long:  ``,
	Run: func(cmd *cobra.Command, params []string) {
		fmt.Println(controllers.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
