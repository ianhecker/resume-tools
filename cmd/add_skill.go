package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/ianhecker/resume-tools/internal/skill"
)

var jsonFile string
var name string
var levelVar string
var yearsVar float64

var addSkillCmd = &cobra.Command{
	Use:   "add-skill",
	Short: "Add a new skill to your JSON file",
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
	addSkillCmd.Flags().StringVarP(&jsonFile, "json", "j", "", "path to the JSON file")

	addSkillCmd.Flags().StringVarP(&name, "name", "n", "", "name of skill")
	addSkillCmd.Flags().Float64VarP(&yearsVar, "years", "y", 0, "years of experience")
	addSkillCmd.Flags().StringVarP(&levelVar, "level", "l", "novice", "experience level (Novice, Intermediate, Advanced, Expert)")

	addSkillCmd.MarkFlagRequired("json")
	addSkillCmd.MarkFlagRequired("name")
}

func readFile(filename string) ([]byte, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		err = fmt.Errorf("could not read file: %w", err)
		checkErr(err)
	}
	return bytes, nil
}

func writeFile(filename string, bytes []byte) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open existing file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(string(bytes))
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}
