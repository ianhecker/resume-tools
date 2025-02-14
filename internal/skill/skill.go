package skill

import (
	"encoding/json"
	"fmt"

	"github.com/ianhecker/resume-tools/internal/experience"
)

type Skill string

type Skills map[string]experience.Experience

func (skills Skills) AddFromRaw(name string, l string, y float64) error {
	experience, err := experience.MakeExperienceFromRaw(l, y)
	if err != nil {
		return err
	}
	skills.Add(name, experience)
	return nil
}

func (skills Skills) Add(name string, experience experience.Experience) {
	skills[name] = experience
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
