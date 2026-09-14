package getfileinfo

import (
	"fmt"

	"github.com/sudzekai-web-os/abstractions"
	"github.com/sudzekai-web-os/files-module/internal/objects/file"
	"github.com/sudzekai-web-os/mediator"
)

type Handler struct {
	logger   abstractions.ILogger
	executor abstractions.IExecutor
}

func NewHandler(
	loggerFactory abstractions.ILoggerFactory,
	executor abstractions.IExecutor,
) mediator.IHandler[Query, QueryResult] {
	logger := loggerFactory.NewLogger("queries:get-file-info")
	return &Handler{
		logger:   logger,
		executor: executor,
	}
}

func (h *Handler) Handle(query Query) QueryResult {
	command := "stat"

	args := []string{
		query.FilePath,
		"-c",
		getCommandFullPattern(),
	}

	result := h.executor.Execute(command, args...)

	if result.Error != nil {
		if result.Stderr != "" {
			return QueryResult{
				Error: fmt.Errorf("%s", result.Stderr),
			}
		}

		return QueryResult{
			Error: result.Error,
		}
	}

	file := file.NewFileInfo(result.Stdout)

	return QueryResult{
		FileInfo: file,
		Error:    nil,
	}
}
