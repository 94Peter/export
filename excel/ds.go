package excel

type DS interface {
	NextPage() (Page, bool)
}

type Page interface {
	GetName() string
	Next() (int, []string)
}
