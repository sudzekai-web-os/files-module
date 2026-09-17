package module

import (
	"github.com/sudzekai-web-os/abstractions"
	"github.com/sudzekai-web-os/files-module/internal/controllers/fileactionscontroller"
	"github.com/sudzekai-web-os/files-module/internal/controllers/filescontroller"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/copyfile"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/createfile"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/deletefile"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/getfileinfo"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/movefile"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/renamefile"
	"github.com/sudzekai-web-os/files-module/internal/utilities/filesystem"
	"github.com/sudzekai-web-os/mediator"
)

type FilesModule struct {
}

func NewFilesModule() abstractions.IModule {
	return &FilesModule{}
}

func (m *FilesModule) Name() string {
	return "Files Module"
}

func (m *FilesModule) Description() string {
	return ""
}

func (m *FilesModule) Version() string {
	return "v0.6.0"
}

func (m *FilesModule) Initialize(
	registry abstractions.IHandlersRegistry,
	loggerFactory abstractions.ILoggerFactory,
	executor abstractions.IExecutor,
) error {
	fileSystem := filesystem.New(executor)

	configureFilesControllerChain(registry, loggerFactory, fileSystem)
	configureFileActionsControllerChain(registry, loggerFactory, fileSystem)

	return nil
}

func configureFilesControllerChain(
	registry abstractions.IHandlersRegistry,
	loggerFactory abstractions.ILoggerFactory,
	fileSystem *filesystem.FileSystem,
) {
	mediator.RegisterHandler(getfileinfo.NewHandler(loggerFactory, fileSystem))

	filesController := filescontroller.New(loggerFactory)

	filesController.AddRoutes(registry)
}

func configureFileActionsControllerChain(
	registry abstractions.IHandlersRegistry,
	loggerFactory abstractions.ILoggerFactory,
	fileSystem *filesystem.FileSystem,
) {
	mediator.RegisterHandler(createfile.NewHandler(loggerFactory, fileSystem))
	mediator.RegisterHandler(deletefile.NewHandler(loggerFactory, fileSystem))
	mediator.RegisterHandler(copyfile.NewHandler(loggerFactory, fileSystem))
	mediator.RegisterHandler(renamefile.NewHandler(loggerFactory, fileSystem))
	mediator.RegisterHandler(movefile.NewHandler(loggerFactory, fileSystem))

	fileActionsController := fileactionscontroller.New(loggerFactory)

	fileActionsController.AddRoutes(registry)
}
