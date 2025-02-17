package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ianhecker/resume-tools/internal/experience"
	"github.com/ianhecker/resume-tools/internal/skill"
)

var (
	skillsFilename  string
	skillName       string
	skillExperience string
	skillYears      float64
)

var skillAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new skill to an existing file",
	Run: func(cmd *cobra.Command, args []string) {
		bytes, err := readFile(skillsFilename)
		checkErr(err)

		skills := skill.MakeSkills()
		err = skills.UnmarshalJSON(bytes)
		checkErr(err)

		skills.AddFromRaw(skillName, skillExperience, skillYears)

		bytes, err = skills.MarshalJSON()
		checkErr(err)

		err = writeFile(skillsFilename, bytes)
		checkErr(err)
	},
}

func init() {
	skillAddCmd.Flags().StringVarP(&skillsFilename, "json", "j", skillsJSON, "path to the JSON file")

	skillAddCmd.Flags().StringVarP(&skillName, "name", "n", "", "name of skill")
	skillAddCmd.Flags().Float64VarP(&skillYears, "years", "y", 0, "years of experience")
	skillAddCmd.Flags().StringVarP(&skillExperience, "experience", "e", experience.Novice.String(), "experience (Novice, Intermediate, Advanced, Expert)")

	skillAddCmd.MarkFlagRequired("name")
}
