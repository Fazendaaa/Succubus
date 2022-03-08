package cmd

import (
	samael "github.com/Fazendaaa/Samael/pkg"
	"github.com/Fazendaaa/Succubus/controllers"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm [package(s) to add]",
	Short: "Removes a packages/libraries from the current project",
	Long:  ``,
	Args:  samael.ValidateProjectPath(projectPath),
	Run:   samael.RunCmd(&projectPath, " removing...", controllers.Rm),
}

func init() {
	rmCmd.Flags().StringVarP(&projectPath, "project path", "p", "", "optional succubus.yaml path remove package from")
	rootCmd.AddCommand(rmCmd)
}
