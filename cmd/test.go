package cmd

import (
	"fmt"

	samael "github.com/Fazendaaa/Samael/pkg"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test [project path]",
	Short: "Tests the current project",
	Long:  ``,
	Args:  samael.ValidateProjectPath(projectPath),
	Run: func(cmd *cobra.Command, params []string) {
		spinner, fail := samael.CreateSpinner(" removing", "")

		if nil != fail {
			fmt.Println()
			fmt.Println(fail)

			return
		}

		// resultChannel := controllers.AddPackages(params, projectPath)
		// fail = consumeChannel(params, "installing", spinner, resultChannel)

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
	testCmd.Flags().StringVarP(&projectPath, "project path", "p", "", "optional succubus.yaml to test project")
	rootCmd.AddCommand(testCmd)
}
