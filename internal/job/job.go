package job

import (
	"encoding/json"
	"fmt"
)

type Jobs map[Title]Experience

func (jobs Jobs) Add(title Title, experience string) {
	experiences := jobs[title]
	if experiences == nil {
		experiences = MakeExperience()
	}

	(&experiences).Add(experience)
	jobs[title] = experiences
}

func (jobs Jobs) MarshalJSON() ([]byte, error) {
	type Alias Jobs

	bytes, err := json.Marshal((*Alias)(&jobs))
	if err != nil {
		return nil, fmt.Errorf("could not marshal jobs: %w", err)
	}
	return bytes, nil
}

func (jobs *Jobs) UnmarshalJSON(data []byte) error {
	type Alias Jobs
	var alias Alias

	err := json.Unmarshal(data, &alias)
	if err != nil {
		return fmt.Errorf("could not unmarshal jobs: %w", err)
	}
	*jobs = Jobs(alias)
	return nil
}
