package job

type Experience []string

func MakeExperience(experiences ...string) Experience {
	return Experience(experiences)
}

func (e *Experience) Add(experience string) {
	*e = append(*e, experience)
}
