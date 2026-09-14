package queries

import "github.com/sudzekai-web-os/files-module/internal/objects/file"

type GetFileInfoQuery struct {
	FilePath string
}

func NewGetFileInfoQuery(filePath string) GetFileInfoQuery {
	return GetFileInfoQuery{
		FilePath: filePath,
	}
}

type GetFileInfoQueryResult struct {
	FileInfo *file.FileInfo
	Error    error
}
