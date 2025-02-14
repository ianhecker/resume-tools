package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Execute() {
	err := rootCmd.Execute()
	checkErr(err)
}

func init() {
	rootCmd.AddCommand(addSkillCmd)
}
