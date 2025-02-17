package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ianhecker/resume-tools/internal/company"
	"github.com/ianhecker/resume-tools/internal/employment"
	"github.com/ianhecker/resume-tools/internal/job"
)

var employmentTemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Generate a fresh template file: " + defaultEmploymentFilename,
	Run: func(cmd *cobra.Command, args []string) {

		company := company.MakeCompany("Company Inc")

		experience := []string{"Increased sales by 35%"}
		job := job.MakeJob("Manager", experience)

		employment := employment.MakeEmployment()
		employment.AddJobForCompany(job, company)

		bytes, err := employment.MarshalJSON()
		checkErr(err)

		err = createFile(defaultEmploymentFilename, bytes)
		checkErr(err)
	},
}
