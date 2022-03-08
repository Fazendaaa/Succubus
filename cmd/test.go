package cmd

import (
	samael "github.com/Fazendaaa/Samael/pkg"
	"github.com/Fazendaaa/Succubus/controllers"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test [project path]",
	Short: "Tests the current project",
	Long:  ``,
	Args:  samael.ValidateProjectPath(projectPath),
	Run:   samael.RunCmd(&projectPath, " testing...", controllers.Test),
}

func init() {
	testCmd.Flags().StringVarP(&projectPath, "project path", "p", "", "optional succubus.yaml to test project")
	rootCmd.AddCommand(testCmd)
}
