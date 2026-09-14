package module

import (
	"github.com/sudzekai-web-os/abstractions"
	"github.com/sudzekai-web-os/files-module/internal/controllers/filescontroller"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/getfileinfo"
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
	return "v0.1.0"
}

func (m *FilesModule) Initialize(
	registry abstractions.IHandlersRegistry,
	loggerFactory abstractions.ILoggerFactory,
	executor abstractions.IExecutor,
) error {
	configureFilesControllerChain(registry, loggerFactory, executor)

	return nil
}

func configureFilesControllerChain(
	registry abstractions.IHandlersRegistry,
	loggerFactory abstractions.ILoggerFactory,
	executor abstractions.IExecutor,
) {
	mediator.RegisterHandler(getfileinfo.NewHandler(loggerFactory, executor))

	filesController := filescontroller.New(loggerFactory)

	filesController.AddRoutes(registry)
}
