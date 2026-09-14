package file

import (
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/sudzekai-web-os/files-module/internal/objects/file/filepermission"
	"github.com/sudzekai-web-os/files-module/internal/objects/file/filetype"
	"github.com/sudzekai-web-os/files-module/internal/objects/statcommand/statformat"
)

type FileInfo struct {
	Name     string
	FullName string

	Type filetype.FileType

	Size      int64
	Inode     uint64
	Links     uint64
	DeviceID  uint64
	RDevice   uint64
	BlockSize int64
	Blocks    int64

	Owner   string
	OwnerID uint32

	Group   string
	GroupID uint32

	LinkTarget string

	SpecialPermissions []filepermission.FilePermission
	UserPermissions    []filepermission.FilePermission
	GroupPermissions   []filepermission.FilePermission
	OthersPermissions  []filepermission.FilePermission

	BirthDateTime        time.Time
	ModificationDateTime time.Time
	ChangeDateTime       time.Time
	AccessDateTime       time.Time

	stats map[string]string
}

func NewFileInfo(stat string) *FileInfo {
	fileInfo := FileInfo{
		stats: make(map[string]string),
	}

	fileInfo.mapStat(stat)
	fileInfo.construct()

	return &fileInfo
}

func (fi *FileInfo) mapStat(stat string) {
	for entry := range strings.SplitSeq(stat, "\n") {
		if idx := strings.Index(entry, ":"); idx != -1 {
			key := strings.TrimSpace(entry[:idx])
			value := strings.TrimSpace(entry[idx+1:])
			fi.stats[key] = value
		}
	}
}

func (fi *FileInfo) construct() {
	fi.Name = fi.getName()
	fi.FullName = fi.getFullName()

	fi.Type = fi.getType()

	fi.Size = fi.getSize()
	fi.Inode = fi.getInode()
	fi.Links = fi.getLinks()
	fi.DeviceID = fi.getDeviceId()
	fi.RDevice = fi.getRDevice()
	fi.BlockSize = fi.getBlockSize()
	fi.Blocks = fi.getBlocks()

	fi.Owner = fi.getOwner()
	fi.OwnerID = fi.getOwnerID()

	fi.Group = fi.getGroup()
	fi.GroupID = fi.getGroupID()

	fi.LinkTarget = fi.getLinkTarget()

	permissions := parsePermissions(fi.getPermissions())

	fi.SpecialPermissions = fi.getSpecialPermissions(permissions)
	fi.UserPermissions = fi.getUserPermissions(permissions)
	fi.GroupPermissions = fi.getGroupPermissions(permissions)
	fi.OthersPermissions = fi.getOthersPermissions(permissions)

	fi.BirthDateTime = fi.getBirthDateTime()
	fi.ModificationDateTime = fi.getModificationDateTime()
	fi.ChangeDateTime = fi.getChangeDateTime()
	fi.AccessDateTime = fi.getAccessDateTime()
}

func (fi *FileInfo) getName() string {
	return filepath.Base(fi.stats[statformat.Name.Key])
}

// подразумевается что передаваемый в stat параметр изначально имеет полный путь к файлу
// что подтверждается валидацией query параметра fullPath у метода GetFileInfo контроллера FilesController
// в связи с чем FullName - результат вывода замены паттерна %n, а name - название конкретного файла
func (fi *FileInfo) getFullName() string {
	return fi.stats[statformat.Name.Key]
}

func (fi *FileInfo) getType() filetype.FileType {
	typeStr := fi.stats["Type"]

	if typeStr == "" {
		return filetype.Unknown
	}

	switch typeStr {
	case "regular file":
		return filetype.Regular
	case "directory":
		return filetype.Directory
	case "symbolic link":
		return filetype.Symlink
	case "socket":
		return filetype.Socket
	case "fifo":
		return filetype.Pipe
	case "block special file":
		return filetype.BlockDevice
	case "character special file":
		return filetype.CharacterDevice
	default:
		return filetype.Unknown
	}
}

func (fi *FileInfo) getSize() int64 {
	sizeStr := fi.stats[statformat.Size.Key]
	size, _ := strconv.ParseInt(sizeStr, 10, 64)
	return size
}

func (fi *FileInfo) getInode() uint64 {
	inodeStr := fi.stats[statformat.Inode.Key]
	inode, _ := strconv.ParseUint(inodeStr, 10, 64)
	return inode
}

func (fi *FileInfo) getLinks() uint64 {
	linksStr := fi.stats[statformat.Links.Key]
	links, _ := strconv.ParseUint(linksStr, 10, 64)
	return links
}

func (fi *FileInfo) getDeviceId() uint64 {
	deviceIDStr := fi.stats[statformat.DeviceID.Key]
	deviceID, _ := strconv.ParseUint(deviceIDStr, 10, 64)
	return deviceID
}

func (fi *FileInfo) getRDevice() uint64 {
	rDeviceStr := fi.stats[statformat.RDevice.Key]
	rDevice, _ := strconv.ParseUint(rDeviceStr, 10, 64)
	return rDevice
}

func (fi *FileInfo) getBlockSize() int64 {
	blockSizeStr := fi.stats[statformat.BlockSize.Key]
	blockSize, _ := strconv.ParseInt(blockSizeStr, 10, 64)
	return blockSize
}

func (fi *FileInfo) getBlocks() int64 {
	blocksStr := fi.stats[statformat.Blocks.Key]
	blocks, _ := strconv.ParseInt(blocksStr, 10, 64)
	return blocks
}

func (fi *FileInfo) getOwner() string {
	return fi.stats[statformat.Owner.Key]
}

func (fi *FileInfo) getOwnerID() uint32 {
	ownerIDStr := fi.stats[statformat.OwnerID.Key]
	ownerID, _ := strconv.ParseUint(ownerIDStr, 10, 32)
	return uint32(ownerID)
}

func (fi *FileInfo) getGroup() string {
	return fi.stats[statformat.Group.Key]
}

func (fi *FileInfo) getGroupID() uint32 {
	groupIDStr := fi.stats[statformat.GroupID.Key]
	groupID, _ := strconv.ParseUint(groupIDStr, 10, 32)
	return uint32(groupID)
}

func (fi *FileInfo) getLinkTarget() string {
	return fi.stats[statformat.LinkTarget.Key]
}

func (fi *FileInfo) getSpecialPermissions(permissions map[string][]filepermission.FilePermission) []filepermission.FilePermission {
	return permissions[statformat.SpecialPermissions.Key]
}

func (fi *FileInfo) getUserPermissions(permissions map[string][]filepermission.FilePermission) []filepermission.FilePermission {
	return permissions[statformat.UserPermissions.Key]
}

func (fi *FileInfo) getGroupPermissions(permissions map[string][]filepermission.FilePermission) []filepermission.FilePermission {
	return permissions[statformat.GroupPermissions.Key]
}

func (fi *FileInfo) getOthersPermissions(permissions map[string][]filepermission.FilePermission) []filepermission.FilePermission {
	return permissions[statformat.OthersPermissions.Key]
}

func (fi *FileInfo) getBirthDateTime() time.Time {
	return parseUnixString(fi.stats[statformat.BirthDateTime.Key])
}

func (fi *FileInfo) getModificationDateTime() time.Time {
	return parseUnixString(fi.stats[statformat.ModificationDateTime.Key])
}

func (fi *FileInfo) getChangeDateTime() time.Time {
	return parseUnixString(fi.stats[statformat.ChangeDateTime.Key])
}

func (fi *FileInfo) getAccessDateTime() time.Time {
	return parseUnixString(fi.stats[statformat.AccessDateTime.Key])
}

func (fi *FileInfo) getPermissions() string {
	return fi.stats[statformat.Permissions.Key]
}

func parsePermissions(permissionsStr string) map[string][]filepermission.FilePermission {
	result := make(map[string][]filepermission.FilePermission)

	if permissionsStr == "" {
		return result
	}

	chars := []rune(permissionsStr)

	if len(chars) < 4 {
		chars = append([]rune{'0'}, chars...)
	}

	for i := 0; i < 4; i++ {
		perms := make([]filepermission.FilePermission, 0, 3)

		if chars[i] == '0' {
			continue
		}

		perm := int(chars[i] - '0')

		if i != 0 {
			if perm&4 != 0 {
				perms = append(perms, filepermission.Read)
			}
			if perm&2 != 0 {
				perms = append(perms, filepermission.Write)
			}
			if perm&1 != 0 {
				perms = append(perms, filepermission.Execute)
			}
		} else {
			if perm&4 != 0 {
				perms = append(perms, filepermission.Setuid)
			}

			if perm&2 != 0 {
				perms = append(perms, filepermission.Setgid)
			}

			if perm&1 != 0 {
				perms = append(perms, filepermission.Sticky)
			}
		}

		keys := []string{
			statformat.SpecialPermissions.Key,
			statformat.UserPermissions.Key,
			statformat.GroupPermissions.Key,
			statformat.OthersPermissions.Key,
		}

		result[keys[i]] = perms
	}

	return result
}

func parseUnixString(unixStr string) time.Time {
	timestamp, _ := strconv.ParseInt(unixStr, 10, 64)

	t := time.Unix(timestamp, 0)
	return t
}
