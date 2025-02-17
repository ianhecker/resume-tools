package company

type Company string

func MakeCompany(name string) Company {
	return Company(name)
}
