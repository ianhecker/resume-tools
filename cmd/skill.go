package cmd

import "github.com/spf13/cobra"

var skillCmd = &cobra.Command{
	Use:   "skill",
	Short: "Generate a list of skills in JSON",
}

func init() {
	skillCmd.AddCommand(skillAddCmd)
	skillCmd.AddCommand(skillTemplateCmd)
}
