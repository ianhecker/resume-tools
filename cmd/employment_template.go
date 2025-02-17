package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ianhecker/resume-tools/internal/company"
	"github.com/ianhecker/resume-tools/internal/employment"
	"github.com/ianhecker/resume-tools/internal/job"
)

var employmentTemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Generate a fresh template file: " + employmentJSON,
	Run: func(cmd *cobra.Command, args []string) {

		company := company.MakeCompany("Company Inc")
		title := job.MakeTitle("Manager")
		experience := "Increased sales by 35%"

		employment := employment.MakeEmployment()
		employment.Add(company, title, experience)

		bytes, err := employment.MarshalJSON()
		checkErr(err)

		err = createFile(employmentJSON, bytes)
		checkErr(err)
	},
}
