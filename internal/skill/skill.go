package skill

import (
	"encoding/json"
	"fmt"

	"github.com/ianhecker/resume-tools/internal/experience"
	"github.com/ianhecker/resume-tools/internal/years"
)

type SkillParameters struct {
	Experience experience.Experience `json:"experience"`
	Years      years.Years           `json:"years"`
}

type Skills map[string]SkillParameters

func MakeSkills() Skills {
	return make(Skills)
}

func (skills Skills) AddFromRaw(name string, e string, y float64) error {
	experience, err := experience.MakeExperienceFromString(e)
	if err != nil {
		return err
	}
	years := years.MakeYears(y)

	skills.Add(name, experience, years)
	return nil
}

func (skills Skills) Add(name string, e experience.Experience, y years.Years) {
	params := SkillParameters{Experience: e, Years: y}
	skills[name] = params
}

func (skills *Skills) MarshalJSON() ([]byte, error) {
	type Alias Skills

	bytes, err := json.Marshal((*Alias)(skills))
	if err != nil {
		return nil, fmt.Errorf("could not marshal skills: %w", err)
	}
	return bytes, nil
}

func (skills *Skills) UnmarshalJSON(data []byte) error {
	type Alias Skills
	var alias Alias

	err := json.Unmarshal(data, &alias)
	if err != nil {
		return fmt.Errorf("could not unmarshal skills: %w", err)
	}
	*skills = Skills(alias)
	return nil
}
