package cmd

import (
	"fmt"

	samael "github.com/Fazendaaa/Samael/pkg"
	"github.com/Fazendaaa/Succubus/controllers"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [package(s) to add]",
	Short: "Adds a packages/libraries to the current project",
	Long:  ``,
	Args:  samael.ValidateProjectPath(projectPath),
	Run: func(cmd *cobra.Command, params []string) {
		spinner, fail := samael.CreateSpinner(" adding", "")

		if nil != fail {
			fmt.Println()
			fmt.Println(fail)

			return
		}

		resultChannel := controllers.Add(params, projectPath)
		fail = samael.ConsumeChannel(params, spinner, resultChannel)

		if nil != fail {
			fmt.Println()
			fmt.Println(fail)

			samael.KillSpinner(spinner, false)

			return
		}

		samael.KillSpinner(spinner, true)
	},
}

func init() {
	addCmd.Flags().StringVarP(&projectPath, "project path", "p", "", "optional succubus.yaml path to add package to")
	rootCmd.AddCommand(addCmd)
}
