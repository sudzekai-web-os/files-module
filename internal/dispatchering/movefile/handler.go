package movefile

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
) mediator.IHandler[Command, CommandResult] {
	return Handler{
		logger:     loggerFactory.NewLogger("handler:move-file"),
		fileSystem: fileSystem,
	}
}

func (hnd Handler) Handle(cmd Command) CommandResult {
	err := hnd.fileSystem.Mv(cmd.FilePath, cmd.DestinationPath)
	return NewResult(err)
}
