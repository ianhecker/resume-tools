package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ianhecker/resume-tools/internal/skill"
)

var jsonFile string
var name string
var levelVar string
var yearsVar float64

var skillAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new skill to an existing file",
	Run: func(cmd *cobra.Command, args []string) {
		bytes, err := readFile(jsonFile)
		checkErr(err)

		var skills = make(skill.Skills)
		err = skills.UnmarshalJSON(bytes)
		checkErr(err)

		skills.AddFromRaw(name, levelVar, yearsVar)

		bytes, err = skills.MarshalJSON()
		checkErr(err)

		err = writeFile(jsonFile, bytes)
		checkErr(err)
	},
}

func init() {
	skillAddCmd.Flags().StringVarP(&jsonFile, "json", "j", "", "path to the JSON file")

	skillAddCmd.Flags().StringVarP(&name, "name", "n", "", "name of skill")
	skillAddCmd.Flags().Float64VarP(&yearsVar, "years", "y", 0, "years of experience")
	skillAddCmd.Flags().StringVarP(&levelVar, "level", "l", "novice", "experience level (Novice, Intermediate, Advanced, Expert)")

	skillAddCmd.MarkFlagRequired("json")
	skillAddCmd.MarkFlagRequired("name")
}
