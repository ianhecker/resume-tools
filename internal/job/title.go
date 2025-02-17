package job

type Title string

func MakeTitle(s string) Title {
	return Title(s)
}
