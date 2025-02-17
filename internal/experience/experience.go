package experience

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Experience int

const (
	Novice Experience = iota + 1
	Intermediate
	Advanced
	Expert
)

func (exp Experience) String() string {
	switch exp {
	case Novice:
		return "novice"
	case Intermediate:
		return "intermediate"
	case Advanced:
		return "advanced"
	case Expert:
		return "expert"
	default:
		return "unknown experience"
	}
}

func ParseExperience(s string) (Experience, error) {
	switch strings.ToLower(s) {
	case "novice":
		return Novice, nil
	case "intermediate":
		return Intermediate, nil
	case "advanced":
		return Advanced, nil
	case "expert":
		return Expert, nil
	default:
		return Experience(0), fmt.Errorf("unknown experience: %s", s)
	}
}

func MakeExperienceFromString(s string) (Experience, error) {
	experience, err := ParseExperience(s)
	if err != nil {
		return 0, fmt.Errorf("could not make experience: %w", err)
	}
	return experience, nil
}

func MakeExperience(experience int) Experience {
	return Experience(experience)
}

func (experience Experience) MarshalJSON() ([]byte, error) {
	bytes, err := json.Marshal(experience.String())
	if err != nil {
		return nil, fmt.Errorf("could not marshal experience: %w", err)
	}
	return bytes, nil
}

func (experience *Experience) UnmarshalJSON(bytes []byte) error {
	var s string
	if err := json.Unmarshal(bytes, &s); err != nil {
		return fmt.Errorf("could not unmarshal experience: %w", err)
	}

	l, err := ParseExperience(s)
	if err != nil {
		return fmt.Errorf("could not parse experience: %w", err)
	}

	*experience = l
	return nil
}
