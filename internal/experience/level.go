package experience

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Level int

const (
	Novice Level = iota + 1
	Intermediate
	Advanced
	Expert
)

func (level Level) String() string {
	switch level {
	case Novice:
		return "novice"
	case Intermediate:
		return "intermediate"
	case Advanced:
		return "advanced"
	case Expert:
		return "expert"
	default:
		return "unknown level"
	}
}

func ParseLevel(s string) (Level, error) {
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
		return Level(0), fmt.Errorf("unknown level: %s", s)
	}
}

func MakeLevelFromString(s string) (Level, error) {
	level, err := ParseLevel(s)
	if err != nil {
		return 0, fmt.Errorf("could not make level: %w", err)
	}
	return level, nil
}

func MakeLevel(level int) Level {
	return Level(level)
}

func (level Level) MarshalJSON() ([]byte, error) {
	bytes, err := json.Marshal(level.String())
	if err != nil {
		return nil, fmt.Errorf("could not marshal level: %w", err)
	}
	return bytes, nil
}

func (level *Level) UnmarshalJSON(bytes []byte) error {
	var s string
	if err := json.Unmarshal(bytes, &s); err != nil {
		return fmt.Errorf("could not unmarshal level: %w", err)
	}

	l, err := ParseLevel(s)
	if err != nil {
		return fmt.Errorf("could not parse level: %w", err)
	}

	*level = l
	return nil
}
