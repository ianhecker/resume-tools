package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ianhecker/resume-tools/internal/skill"
)

var jsonFile string
var name string
var experienceVar string
var yearsVar float64

var skillAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new skill to an existing file",
	Run: func(cmd *cobra.Command, args []string) {
		bytes, err := readFile(jsonFile)
		checkErr(err)

		skills := skill.MakeSkills()
		err = skills.UnmarshalJSON(bytes)
		checkErr(err)

		skills.AddFromRaw(name, experienceVar, yearsVar)

		bytes, err = skills.MarshalJSON()
		checkErr(err)

		err = writeFile(jsonFile, bytes)
		checkErr(err)
	},
}

func init() {
	skillAddCmd.Flags().StringVarP(&jsonFile, "json", "j", defaultSkillsFilename, "path to the JSON file")

	skillAddCmd.Flags().StringVarP(&name, "name", "n", "", "name of skill")
	skillAddCmd.Flags().Float64VarP(&yearsVar, "years", "y", 0, "years of experience")
	skillAddCmd.Flags().StringVarP(&experienceVar, "experience", "e", "novice", "experience (Novice, Intermediate, Advanced, Expert)")

	skillAddCmd.MarkFlagRequired("name")
}
