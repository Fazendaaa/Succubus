package cmd

import (
	"fmt"

	samael "github.com/Fazendaaa/Samael/pkg"
	"github.com/Fazendaaa/Succubus/controllers"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [project path]",
	Short: "Runs the current project",
	Long:  ``,
	Args:  samael.ValidateProjectPath(projectPath),
	Run: func(cmd *cobra.Command, params []string) {
		spinner, fail := samael.CreateSpinner(" removing", "")

		if nil != fail {
			fmt.Println()
			fmt.Println(fail)

			return
		}

		resultChannel := controllers.Run(params, projectPath)
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
	runCmd.Flags().StringVarP(&projectPath, "project path", "p", "", "optional succubus.yaml to run project")
	rootCmd.AddCommand(runCmd)
}
