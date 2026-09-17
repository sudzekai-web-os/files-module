package readfile

type Query struct {
	FilePath string
}

func NewQuery(filePath string) Query {
	return Query{
		FilePath: filePath,
	}
}
