package job

type Job struct {
	Title      string   `json:"title"`
	Experience []string `json:"experience"`
}

func MakeJob(title string, experience []string) Job {
	return Job{Title: title, Experience: experience}
}
