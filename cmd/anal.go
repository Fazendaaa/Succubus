package cmd

import (
	samael "github.com/Fazendaaa/Samael/pkg"
	"github.com/Fazendaaa/Succubus/controllers"
	"github.com/spf13/cobra"
)

var analCmd = &cobra.Command{
	Use:   "anal [project path]",
	Short: "Runs code analysis in the current project",
	Long:  ``,
	Args:  samael.ValidateProjectPath(projectPath),
	Run:   samael.RunCmd(&projectPath, " analysis...", controllers.Anal),
}

func init() {
	analCmd.Flags().StringVarP(&projectPath, "project path", "p", "", "optional succubus.yaml to run code analysis")
	rootCmd.AddCommand(analCmd)
}
