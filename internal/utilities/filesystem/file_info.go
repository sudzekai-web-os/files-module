package filesystem

import (
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type FileInfo struct {
	Name     string
	FullName string

	Type FileType

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

	SpecialPermissions []FilePermission
	UserPermissions    []FilePermission
	GroupPermissions   []FilePermission
	OthersPermissions  []FilePermission

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
	return filepath.Base(fi.stats["Name"])
}

// подразумевается что передаваемый в stat параметр изначально имеет полный путь к файлу
// что подтверждается валидацией query параметра fullPath у метода GetFileInfo контроллера FilesController
// в связи с чем FullName - результат вывода замены паттерна %n, а name - название конкретного файла
func (fi *FileInfo) getFullName() string {
	return fi.stats["Name"]
}

func (fi *FileInfo) getType() FileType {
	typeStr := fi.stats["Type"]

	if typeStr == "" {
		return TypeUnknown
	}

	switch typeStr {
	case "regular file", "regular empty file":
		return TypeRegular
	case "directory":
		return TypeDirectory
	case "symbolic link":
		return TypeSymlink
	case "socket":
		return TypeSocket
	case "fifo":
		return TypePipe
	case "block special file":
		return TypeBlockDevice
	case "character special file":
		return TypeCharacterDevice
	default:
		return TypeUnknown
	}
}

func (fi *FileInfo) getSize() int64 {
	sizeStr := fi.stats["Size"]
	size, _ := strconv.ParseInt(sizeStr, 10, 64)
	return size
}

func (fi *FileInfo) getInode() uint64 {
	inodeStr := fi.stats["Inode"]
	inode, _ := strconv.ParseUint(inodeStr, 10, 64)
	return inode
}

func (fi *FileInfo) getLinks() uint64 {
	linksStr := fi.stats["Links"]
	links, _ := strconv.ParseUint(linksStr, 10, 64)
	return links
}

func (fi *FileInfo) getDeviceId() uint64 {
	deviceIDStr := fi.stats["DeviceID"]
	deviceID, _ := strconv.ParseUint(deviceIDStr, 10, 64)
	return deviceID
}

func (fi *FileInfo) getRDevice() uint64 {
	rDeviceStr := fi.stats["RDevice"]
	rDevice, _ := strconv.ParseUint(rDeviceStr, 10, 64)
	return rDevice
}

func (fi *FileInfo) getBlockSize() int64 {
	blockSizeStr := fi.stats["BlockSize"]
	blockSize, _ := strconv.ParseInt(blockSizeStr, 10, 64)
	return blockSize
}

func (fi *FileInfo) getBlocks() int64 {
	blocksStr := fi.stats["Blocks"]
	blocks, _ := strconv.ParseInt(blocksStr, 10, 64)
	return blocks
}

func (fi *FileInfo) getOwner() string {
	return fi.stats["Owner"]
}

func (fi *FileInfo) getOwnerID() uint32 {
	ownerIDStr := fi.stats["OwnerID"]
	ownerID, _ := strconv.ParseUint(ownerIDStr, 10, 32)
	return uint32(ownerID)
}

func (fi *FileInfo) getGroup() string {
	return fi.stats["Group"]
}

func (fi *FileInfo) getGroupID() uint32 {
	groupIDStr := fi.stats["GroupID"]
	groupID, _ := strconv.ParseUint(groupIDStr, 10, 32)
	return uint32(groupID)
}

func (fi *FileInfo) getLinkTarget() string {
	return fi.stats["LinkTarget"]
}

func (fi *FileInfo) getSpecialPermissions(permissions map[string][]FilePermission) []FilePermission {
	return permissions["SpecialPermissions"]
}

func (fi *FileInfo) getUserPermissions(permissions map[string][]FilePermission) []FilePermission {
	return permissions["UserPermissions"]
}

func (fi *FileInfo) getGroupPermissions(permissions map[string][]FilePermission) []FilePermission {
	return permissions["GroupPermissions"]
}

func (fi *FileInfo) getOthersPermissions(permissions map[string][]FilePermission) []FilePermission {
	return permissions["OthersPermissions"]
}

func (fi *FileInfo) getBirthDateTime() time.Time {
	return parseUnixString(fi.stats["BirthDateTime"])
}

func (fi *FileInfo) getModificationDateTime() time.Time {
	return parseUnixString(fi.stats["ModificationDateTime"])
}

func (fi *FileInfo) getChangeDateTime() time.Time {
	return parseUnixString(fi.stats["ChangeDateTime"])
}

func (fi *FileInfo) getAccessDateTime() time.Time {
	return parseUnixString(fi.stats["AccessDateTime"])
}

func (fi *FileInfo) getPermissions() string {
	return fi.stats["Permissions"]
}

func parsePermissions(permissionsStr string) map[string][]FilePermission {
	result := make(map[string][]FilePermission)

	if permissionsStr == "" {
		return result
	}

	chars := []rune(permissionsStr)

	if len(chars) < 4 {
		chars = append([]rune{'0'}, chars...)
	}

	for i := 0; i < 4; i++ {
		perms := make([]FilePermission, 0, 3)

		if chars[i] == '0' {
			continue
		}

		perm := int(chars[i] - '0')

		if i != 0 {
			if perm&4 != 0 {
				perms = append(perms, PermissionRead)
			}
			if perm&2 != 0 {
				perms = append(perms, PermissionWrite)
			}
			if perm&1 != 0 {
				perms = append(perms, PermissionExecute)
			}
		} else {
			if perm&4 != 0 {
				perms = append(perms, PermissionSetuid)
			}

			if perm&2 != 0 {
				perms = append(perms, PermissionSetgid)
			}

			if perm&1 != 0 {
				perms = append(perms, PermissionSticky)
			}
		}

		keys := []string{
			"SpecialPermissions",
			"UserPermissions",
			"GroupPermissions",
			"OthersPermissions",
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
