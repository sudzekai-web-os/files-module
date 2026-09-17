package getfileinfo

import "github.com/sudzekai-web-os/files-module/internal/utilities/filesystem"

type QueryResult struct {
	FileInfo *filesystem.FileInfo
	Error    error
}

func NewResult(fi *filesystem.FileInfo, err error) QueryResult {
	return QueryResult{
		FileInfo: fi,
		Error:    err,
	}
}
