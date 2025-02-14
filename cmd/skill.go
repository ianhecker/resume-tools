package cmd

import "github.com/spf13/cobra"

var defaultSkillsFilename string = "skills.json"

var skillCmd = &cobra.Command{
	Use:     "skill",
	Aliases: []string{"skills"},
	Short:   "Generate a list of skills in JSON",
}

func init() {
	skillCmd.AddCommand(skillAddCmd)
	skillCmd.AddCommand(skillTemplateCmd)
}
