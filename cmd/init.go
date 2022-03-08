package cmd

import (
	samael "github.com/Fazendaaa/Samael/pkg"
	"github.com/Fazendaaa/Succubus/controllers"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [project's path]",
	Short: "Initialization of the project project",
	Long:  ``,
	Args:  samael.ValidateProjectPath(projectPath),
	Run:   samael.RunCmd(&projectPath, " initializing...", controllers.Init),
}

func init() {
	initCmd.Flags().StringVarP(&projectPath, "project path", "p", "", "optional succubus.yaml to init project")
	rootCmd.AddCommand(initCmd)
}
