package module

import (
	"github.com/sudzekai-web-os/abstractions"
	queryhandlers "github.com/sudzekai-web-os/files-module/internal/application/query_handlers"
	"github.com/sudzekai-web-os/files-module/internal/presentation/controllers"
	querydispatcher "github.com/sudzekai-web-os/mediator/querydispatcher"
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
	return "v0.0.1"
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
	querydispatcher.RegisterHandler(queryhandlers.NewGetFileInfoQueryHandler(loggerFactory, executor))

	filesController := controllers.NewFilesController(loggerFactory, executor)
	filesController.AddRoutes(registry)
}
