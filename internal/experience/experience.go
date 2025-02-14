package experience

import (
	"fmt"
)

type Experience struct {
	Level Level `json:"level"`
	Years Years `json:"years"`
}

func MakeExperience(level Level, years Years) Experience {
	return Experience{Level: level, Years: years}
}

func MakeExperienceFromRaw(l string, y float64) (Experience, error) {
	level, err := MakeLevelFromString(l)
	if err != nil {
		return Experience{}, fmt.Errorf("could not make experience: %w", err)
	}
	years := MakeYears(y)
	return MakeExperience(level, years), nil
}
