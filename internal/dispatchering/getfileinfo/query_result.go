package getfileinfo

import "github.com/sudzekai-web-os/files-module/internal/objects/file"

type QueryResult struct {
	FileInfo *file.FileInfo
	Error    error
}

func NewResult(fi *file.FileInfo, err error) QueryResult {
	return QueryResult{
		FileInfo: fi,
		Error:    err,
	}
}
