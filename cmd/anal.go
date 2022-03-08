package cmd

import (
	"fmt"

	samael "github.com/Fazendaaa/Samael/pkg"
	"github.com/spf13/cobra"
)

var analCmd = &cobra.Command{
	Use:   "anal [project path]",
	Short: "Runs code analysis in the current project",
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
	analCmd.Flags().StringVarP(&projectPath, "project path", "p", "", "optional succubus.yaml to run code analysis")
	rootCmd.AddCommand(analCmd)
}
