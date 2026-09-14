package queries

type GetFileInfoQuery struct {
	FilePath string
}

func NewGetFileInfoQuery(filePath string) *GetFileInfoQuery {
	return &GetFileInfoQuery{
		FilePath: filePath,
	}
}

func (GetFileInfoQuery) IsQuery() {}
