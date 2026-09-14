package controllers

import (
	"net/http"

	"github.com/sudzekai-web-os/abstractions"
	"github.com/sudzekai-web-os/files-module/internal/application/queries"
	"github.com/sudzekai-web-os/files-module/internal/presentation/interfaces"
	"github.com/sudzekai-web-os/mediator"
	"github.com/sudzekai-web-os/types"
)

type FilesContoller struct {
	loggerFactory abstractions.ILoggerFactory
}

func NewFilesController(
	loggerFactory abstractions.ILoggerFactory,
) interfaces.IFilesController {
	return &FilesContoller{
		loggerFactory: loggerFactory,
	}
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

	var queryResult queries.GetFileInfoQueryResult

	err = mediator.Dispatch(queries.NewGetFileInfoQuery(filePath), &queryResult)

	if err != nil {
		log := fc.loggerFactory.NewLogger("files-module:get-file-info")
		log.LogError("%s", err.Error())

		result.Data = "Внутренняя ошибка сервера"
		result.StatusCode = 500
		return
	}

	result.Data = queryResult
	result.StatusCode = 200

	return
}
