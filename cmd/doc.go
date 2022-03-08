package cmd

import (
	"fmt"

	samael "github.com/Fazendaaa/Samael/pkg"
	"github.com/Fazendaaa/Succubus/controllers"
	"github.com/spf13/cobra"
)

var docCmd = &cobra.Command{
	Use:   "doc [package(s) to add]",
	Short: "Documentation in the current project",
	Long:  ``,
	Args:  samael.ValidateProjectPath(projectPath),
	Run: func(cmd *cobra.Command, params []string) {
		spinner, fail := samael.CreateSpinner(" removing", "")

		if nil != fail {
			fmt.Println()
			fmt.Println(fail)

			return
		}

		resultChannel := controllers.Doc(params, projectPath)
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
	docCmd.Flags().StringVarP(&projectPath, "project path", "p", "", "optional succubus.yaml to run documentation process")
	rootCmd.AddCommand(docCmd)
}
