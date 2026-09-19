package module

import (
	"github.com/sudzekai-web-os/core"
	"github.com/sudzekai-web-os/files-module/internal/controllers/fileactionscontroller"
	"github.com/sudzekai-web-os/files-module/internal/controllers/filecontentactionscontroller"
	"github.com/sudzekai-web-os/files-module/internal/controllers/filescontroller"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/appendfile"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/copyfile"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/createfile"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/deletefile"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/getfileinfo"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/movefile"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/readfile"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/renamefile"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/writefile"
	"github.com/sudzekai-web-os/files-module/internal/utilities/filesystem"
	"github.com/sudzekai-web-os/mediator"
)

type FilesModule struct {
	registry      core.IHandlersRegistry
	loggerFactory core.ILoggerFactory
	executor      core.IExecutor
}

func NewFilesModule() *FilesModule {
	return &FilesModule{}
}

func (m *FilesModule) Name() string {
	return "Files Module"
}

func (m *FilesModule) Description() string {
	return "Модуль предоставляет набор HTTP-обработчиков для управления файлами: получения информации, создания, удаления, копирования, перемещения и переименования. Также модуль поддерживает чтение, полную перезапись и добавление содержимого файлов с обработкой ошибок файловой системы."
}

func (m *FilesModule) Version() string {
	return "v0.9.0"
}

func (m *FilesModule) Start() error {
	fileSystem := filesystem.New(m.executor)

	configureFilesControllerChain(m.registry, m.loggerFactory, fileSystem)
	configureFileActionsControllerChain(m.registry, m.loggerFactory, fileSystem)
	configureFileContentActionsControllerChain(m.registry, m.loggerFactory, fileSystem)

	return nil
}

func (m *FilesModule) AddHandlersRegistry(registry core.IHandlersRegistry) {
	m.registry = registry
}

func (m *FilesModule) AddLoggerFactory(loggerFactory core.ILoggerFactory) {
	m.loggerFactory = loggerFactory
}

func (m *FilesModule) AddExecutor(executor core.IExecutor) {
	m.executor = executor
}

func configureFilesControllerChain(
	registry core.IHandlersRegistry,
	loggerFactory core.ILoggerFactory,
	fileSystem *filesystem.FileSystem,
) {
	mediator.RegisterHandler(getfileinfo.NewHandler(loggerFactory, fileSystem))

	filesController := filescontroller.New(loggerFactory)

	filesController.AddRoutes(registry)
}

func configureFileActionsControllerChain(
	registry core.IHandlersRegistry,
	loggerFactory core.ILoggerFactory,
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

func configureFileContentActionsControllerChain(
	registry core.IHandlersRegistry,
	loggerFactory core.ILoggerFactory,
	fileSystem *filesystem.FileSystem,
) {
	mediator.RegisterHandler(appendfile.NewHandler(loggerFactory, fileSystem))
	mediator.RegisterHandler(writefile.NewHandler(loggerFactory, fileSystem))
	mediator.RegisterHandler(readfile.NewHandler(loggerFactory, fileSystem))

	fileContentActionsController := filecontentactionscontroller.New(loggerFactory)

	fileContentActionsController.AddRoutes(registry)
}
