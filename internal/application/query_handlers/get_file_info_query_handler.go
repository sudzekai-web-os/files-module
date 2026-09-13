package queryhandlers

import (
	"fmt"
	"strings"

	"github.com/sudzekai-web-os/abstractions"
	"github.com/sudzekai-web-os/files-module/internal/application/queries"
	"github.com/sudzekai-web-os/files-module/internal/objects/file"
	"github.com/sudzekai-web-os/files-module/internal/objects/statcommand/statkeypatterns"
	mediator "github.com/sudzekai-web-os/mediator/abstractions"
)

type GetFileInfoQueryHandler struct {
	logger   abstractions.ILogger
	executor abstractions.IExecutor
}

func NewGetFileInfoQueryHandler(
	loggerFactory abstractions.ILoggerFactory,
	executor abstractions.IExecutor,
) mediator.IQueryHandler[queries.GetFileInfoQuery, *file.FileInfo] {
	logger := loggerFactory.NewLogger("queries:get-file-info")
	return &GetFileInfoQueryHandler{
		logger:   logger,
		executor: executor,
	}
}

func (h *GetFileInfoQueryHandler) Handle(query queries.GetFileInfoQuery) (*file.FileInfo, error) {
	command := "stat"
	args := []string{
		query.FilePath,
		"-c",
		getCommandFullPattern(),
	}

	result := h.executor.Execute(command, args...)

	if result.Error != nil {
		return nil, result.Error
	}

	file := file.NewFileInfo(result.Stdout)

	return file, nil
}

func getCommandFullPattern() string {
	patterns := []string{
		getCommandPattern(statkeypatterns.Name.Key, statkeypatterns.Name.Pattern),
		getCommandPattern(statkeypatterns.FullName.Key, statkeypatterns.FullName.Pattern),
		getCommandPattern(statkeypatterns.Type.Key, statkeypatterns.Type.Pattern),
		getCommandPattern(statkeypatterns.Size.Key, statkeypatterns.Size.Pattern),
		getCommandPattern(statkeypatterns.Inode.Key, statkeypatterns.Inode.Pattern),
		getCommandPattern(statkeypatterns.Links.Key, statkeypatterns.Links.Pattern),
		getCommandPattern(statkeypatterns.DeviceID.Key, statkeypatterns.DeviceID.Pattern),
		getCommandPattern(statkeypatterns.RDevice.Key, statkeypatterns.RDevice.Pattern),
		getCommandPattern(statkeypatterns.BlockSize.Key, statkeypatterns.BlockSize.Pattern),
		getCommandPattern(statkeypatterns.Blocks.Key, statkeypatterns.Blocks.Pattern),
		getCommandPattern(statkeypatterns.Owner.Key, statkeypatterns.Owner.Pattern),
		getCommandPattern(statkeypatterns.OwnerID.Key, statkeypatterns.OwnerID.Pattern),
		getCommandPattern(statkeypatterns.Group.Key, statkeypatterns.Group.Pattern),
		getCommandPattern(statkeypatterns.GroupID.Key, statkeypatterns.GroupID.Pattern),
		getCommandPattern(statkeypatterns.LinkTarget.Key, statkeypatterns.LinkTarget.Pattern),
		getCommandPattern(statkeypatterns.Permissions.Key, statkeypatterns.Permissions.Pattern),
		getCommandPattern(statkeypatterns.BirthDateTime.Key, statkeypatterns.BirthDateTime.Pattern),
		getCommandPattern(statkeypatterns.ModificationDateTime.Key, statkeypatterns.ModificationDateTime.Pattern),
		getCommandPattern(statkeypatterns.ChangeDateTime.Key, statkeypatterns.ChangeDateTime.Pattern),
		getCommandPattern(statkeypatterns.AccessDateTime.Key, statkeypatterns.AccessDateTime.Pattern),
	}
	return fmt.Sprintf(
		"$\"%s\"", strings.Join(patterns, "\\n"),
	)
}

func getCommandPattern(key string, pattern string) string {
	return fmt.Sprintf("%s:%s", key, pattern)
}
