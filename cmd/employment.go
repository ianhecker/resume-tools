package cmd

import "github.com/spf13/cobra"

var defaultEmploymentFilename string = "employment.json"

var employmentCmd = &cobra.Command{
	Use:   "employment",
	Short: "Generate a list of employment in JSON",
}

func init() {
	// employmentCmd.AddCommand(employmentAddCmd)
	employmentCmd.AddCommand(employmentTemplateCmd)
}
