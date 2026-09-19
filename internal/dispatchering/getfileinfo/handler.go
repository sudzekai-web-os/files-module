package getfileinfo

import (
	"github.com/sudzekai-web-os/core"
	"github.com/sudzekai-web-os/files-module/internal/utilities/filesystem"
	"github.com/sudzekai-web-os/mediator"
)

type Handler struct {
	logger     core.ILogger
	fileSystem *filesystem.FileSystem
}

func NewHandler(
	loggerFactory core.ILoggerFactory,
	fileSystem *filesystem.FileSystem,
) mediator.IHandler[Query, QueryResult] {
	logger := loggerFactory.NewLogger("handler:get-file-info")
	return &Handler{
		logger:     logger,
		fileSystem: fileSystem,
	}
}

func (h *Handler) Handle(query Query) QueryResult {
	fileInfo, err := h.fileSystem.Stat(query.FilePath)

	return NewResult(
		fileInfo,
		err,
	)
}
