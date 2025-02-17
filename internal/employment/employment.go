package employment

import (
	"encoding/json"
	"fmt"

	"github.com/ianhecker/resume-tools/internal/company"
	"github.com/ianhecker/resume-tools/internal/job"
)

type Employment map[company.Company]job.Jobs

func MakeEmployment() Employment {
	return make(Employment)
}

func (employment Employment) Add(
	company company.Company,
	title job.Title,
	experience string,
) {
	jobs := employment[company]
	if jobs == nil {
		jobs = make(job.Jobs)
	}

	jobs.Add(title, experience)
	employment[company] = jobs
}

func (employment Employment) MarshalJSON() ([]byte, error) {
	type Alias Employment

	bytes, err := json.Marshal((*Alias)(&employment))
	if err != nil {
		return nil, fmt.Errorf("could not marshal employment: %w", err)
	}
	return bytes, nil
}

func (employment *Employment) UnmarshalJSON(data []byte) error {
	type Alias Employment
	var alias Alias

	err := json.Unmarshal(data, &alias)
	if err != nil {
		return fmt.Errorf("could not unmarshal employment: %w", err)
	}
	*employment = Employment(alias)
	return nil
}
