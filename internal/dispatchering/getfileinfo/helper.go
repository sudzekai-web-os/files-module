package getfileinfo

import (
	"fmt"
	"strings"

	"github.com/sudzekai-web-os/files-module/internal/objects/statcommand/statformat"
)

func getCommandFullPattern() string {
	patterns := []string{
		getCommandPattern(statformat.Name),
		getCommandPattern(statformat.FullName),
		getCommandPattern(statformat.Type),
		getCommandPattern(statformat.Size),
		getCommandPattern(statformat.Inode),
		getCommandPattern(statformat.Links),
		getCommandPattern(statformat.DeviceID),
		getCommandPattern(statformat.RDevice),
		getCommandPattern(statformat.BlockSize),
		getCommandPattern(statformat.Blocks),
		getCommandPattern(statformat.Owner),
		getCommandPattern(statformat.OwnerID),
		getCommandPattern(statformat.Group),
		getCommandPattern(statformat.GroupID),
		getCommandPattern(statformat.LinkTarget),
		getCommandPattern(statformat.Permissions),
		getCommandPattern(statformat.BirthDateTime),
		getCommandPattern(statformat.ModificationDateTime),
		getCommandPattern(statformat.ChangeDateTime),
		getCommandPattern(statformat.AccessDateTime),
	}

	return strings.Join(patterns, "\n")
}

func getCommandPattern(format statformat.StatFormat) string {
	return fmt.Sprintf("%s:%s", format.Key, format.Format)
}
