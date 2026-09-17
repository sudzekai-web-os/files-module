package readfile

type QueryResult struct {
	Content string
	Error   error
}

func NewResult(content string, err error) QueryResult {
	return QueryResult{
		Content: content,
		Error:   err,
	}
}
