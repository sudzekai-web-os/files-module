package controllers

import (
	"net/http"

	"github.com/sudzekai-web-os/abstractions"
	"github.com/sudzekai-web-os/files-module/internal/application/queries"
	"github.com/sudzekai-web-os/files-module/internal/objects/file"
	"github.com/sudzekai-web-os/files-module/internal/presentation/interfaces"
	querydispatcher "github.com/sudzekai-web-os/mediator/querydispatcher"
	"github.com/sudzekai-web-os/types"
)

type FilesContoller struct {
}

func NewFilesController(
	loggerFactory abstractions.ILoggerFactory,
	executor abstractions.IExecutor,
) interfaces.IFilesController {
	return &FilesContoller{}
}

func (fc *FilesContoller) AddRoutes(registry abstractions.IHandlersRegistry) {
	registry.AddHandler("GET /files/info", fc.GetFileInfo)
}

func (fc *FilesContoller) GetFileInfo(r *http.Request) (result types.HandlerResult) {
	filePath := r.URL.Query().Get("filePath")

	err := validateFilePath(filePath)

	if err != nil {
		result.StatusCode = 400
		result.Error = err
		return
	}

	queryResult, err := querydispatcher.Query[*queries.GetFileInfoQuery, *file.FileInfo](queries.NewGetFileInfoQuery(filePath))

	result.Data = queryResult
	result.StatusCode = 200

	return
}
