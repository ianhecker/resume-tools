package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ianhecker/resume-tools/internal/experience"
	"github.com/ianhecker/resume-tools/internal/skill"
)

var defaultFilename string = "skills.json"

var skillTemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Generate a fresh template file: " + defaultFilename,
	Run: func(cmd *cobra.Command, args []string) {
		var skills = make(skill.Skills)
		skills.AddFromRaw(name, experience.Level(0).String(), 0)

		bytes, err := skills.MarshalJSON()
		checkErr(err)

		err = writeFile(defaultFilename, bytes)
		checkErr(err)
	},
}
