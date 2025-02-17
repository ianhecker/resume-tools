package cmd

import (
	"github.com/ianhecker/resume-tools/internal/company"
	"github.com/ianhecker/resume-tools/internal/employment"
	"github.com/ianhecker/resume-tools/internal/job"
	"github.com/spf13/cobra"
)

var (
	employmentFilename   string
	employmentCompany    string
	employmentExperience string
	employmentTitle      string
)

var employmentAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add job experience to a new or existing company",
	Run: func(cmd *cobra.Command, args []string) {
		bytes, err := readFile(employmentFilename)
		checkErr(err)

		employment := employment.MakeEmployment()
		err = employment.UnmarshalJSON(bytes)
		checkErr(err)

		company := company.MakeCompany(employmentCompany)
		title := job.MakeTitle(employmentTitle)

		employment.Add(company, title, employmentExperience)

		bytes, err = employment.MarshalJSON()
		checkErr(err)

		err = writeFile(employmentFilename, bytes)
		checkErr(err)
	},
}

func init() {
	employmentAddCmd.Flags().StringVarP(&employmentFilename, "json", "j", employmentJSON, "path to the JSON file")

	employmentAddCmd.Flags().StringVarP(&employmentCompany, "company", "c", "", "company name at employment")
	employmentAddCmd.Flags().StringVarP(&employmentTitle, "title", "t", "", "job title at employment")
	employmentAddCmd.Flags().StringVarP(&employmentExperience, "experience", "e", "", "bullet point of work at employment")

	employmentAddCmd.MarkFlagRequired("company")
	employmentAddCmd.MarkFlagRequired("title")
}
