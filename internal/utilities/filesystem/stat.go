package filesystem

import (
	"strings"
)

func getCommandFullPattern() string {
	patterns := []string{
		"Name:%n",
		"FullName:",
		"Type:%F",
		"Size:%s",
		"Inode:%i",
		"Links:%h",
		"DeviceID:%d",
		"RDevice:%r",
		"BlockSize:%B",
		"Blocks:%b",
		"Owner:%U",
		"OwnerID:%u",
		"Group:%G",
		"GroupID:%g",
		"LinkTarget:%N",
		"Permissions:%a",
		"BirthDateTime:%W",
		"ModificationDateTime:%Y",
		"ChangeDateTime:%Z",
		"AccessDateTime:%X",
	}

	return strings.Join(patterns, "\n")
}
