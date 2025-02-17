package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ianhecker/resume-tools/internal/experience"
	"github.com/ianhecker/resume-tools/internal/skill"
)

var skillTemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Generate a fresh template file: " + skillsJSON,
	Run: func(cmd *cobra.Command, args []string) {

		var skills = make(skill.Skills)
		err := skills.AddFromRaw("underwater basket weaving", experience.Novice.String(), 0)
		checkErr(err)

		bytes, err := skills.MarshalJSON()
		checkErr(err)

		err = createFile(skillsJSON, bytes)
		checkErr(err)
	},
}
