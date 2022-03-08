package cmd

import (
	samael "github.com/Fazendaaa/Samael/pkg"
	"github.com/Fazendaaa/Succubus/controllers"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [package(s) to add]",
	Short: "Adds a packages/libraries to the current project",
	Long:  ``,
	Args:  samael.ValidateProjectPath(projectPath),
	Run:   samael.RunCmd(&projectPath, " adding...", controllers.Add),
}

func init() {
	addCmd.Flags().StringVarP(&projectPath, "project path", "p", "", "optional succubus.yaml path to add package to")
	rootCmd.AddCommand(addCmd)
}
