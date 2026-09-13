package file

import (
	"strconv"
	"strings"
	"time"

	"github.com/sudzekai-web-os/files-module/internal/objects/file/filepermissions"
	"github.com/sudzekai-web-os/files-module/internal/objects/file/filetypes"
	"github.com/sudzekai-web-os/files-module/internal/objects/statcommand/statkeypatterns"
)

type FileInfo struct {
	Name     string
	FullName string

	Type filetypes.FileType

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

	SpecialPermissions []filepermissions.FilePermission
	UserPermissions    []filepermissions.FilePermission
	GroupPermissions   []filepermissions.FilePermission
	OthersPermissions  []filepermissions.FilePermission

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
	return fi.stats["Name"]
}

func (fi *FileInfo) getFullName() string {
	return fi.stats["FullName"]
}

func (fi *FileInfo) getType() filetypes.FileType {
	typeStr := fi.stats["Type"]

	if typeStr == "" {
		return filetypes.Unknown
	}

	switch typeStr {
	case "regular file":
		return filetypes.Regular
	case "directory":
		return filetypes.Directory
	case "symbolic link":
		return filetypes.Symlink
	case "socket":
		return filetypes.Socket
	case "fifo":
		return filetypes.Pipe
	case "block special file":
		return filetypes.BlockDevice
	case "character special file":
		return filetypes.CharacterDevice
	default:
		return filetypes.Unknown
	}
}

func (fi *FileInfo) getSize() int64 {
	sizeStr := fi.stats[statkeypatterns.Size.Key]
	size, _ := strconv.ParseInt(sizeStr, 10, 64)
	return size
}

func (fi *FileInfo) getInode() uint64 {
	inodeStr := fi.stats[statkeypatterns.Inode.Key]
	inode, _ := strconv.ParseUint(inodeStr, 10, 64)
	return inode
}

func (fi *FileInfo) getLinks() uint64 {
	linksStr := fi.stats[statkeypatterns.Links.Key]
	links, _ := strconv.ParseUint(linksStr, 10, 64)
	return links
}

func (fi *FileInfo) getDeviceId() uint64 {
	deviceIDStr := fi.stats[statkeypatterns.DeviceID.Key]
	deviceID, _ := strconv.ParseUint(deviceIDStr, 10, 64)
	return deviceID
}

func (fi *FileInfo) getRDevice() uint64 {
	rDeviceStr := fi.stats[statkeypatterns.RDevice.Key]
	rDevice, _ := strconv.ParseUint(rDeviceStr, 10, 64)
	return rDevice
}

func (fi *FileInfo) getBlockSize() int64 {
	blockSizeStr := fi.stats[statkeypatterns.BlockSize.Key]
	blockSize, _ := strconv.ParseInt(blockSizeStr, 10, 64)
	return blockSize
}

func (fi *FileInfo) getBlocks() int64 {
	blocksStr := fi.stats[statkeypatterns.Blocks.Key]
	blocks, _ := strconv.ParseInt(blocksStr, 10, 64)
	return blocks
}

func (fi *FileInfo) getOwner() string {
	return fi.stats[statkeypatterns.Owner.Key]
}

func (fi *FileInfo) getOwnerID() uint32 {
	ownerIDStr := fi.stats[statkeypatterns.OwnerID.Key]
	ownerID, _ := strconv.ParseUint(ownerIDStr, 10, 32)
	return uint32(ownerID)
}

func (fi *FileInfo) getGroup() string {
	return fi.stats[statkeypatterns.Group.Key]
}

func (fi *FileInfo) getGroupID() uint32 {
	groupIDStr := fi.stats[statkeypatterns.GroupID.Key]
	groupID, _ := strconv.ParseUint(groupIDStr, 10, 32)
	return uint32(groupID)
}

func (fi *FileInfo) getLinkTarget() string {
	return fi.stats[statkeypatterns.LinkTarget.Key]
}

func (fi *FileInfo) getSpecialPermissions(permissions map[string][]filepermissions.FilePermission) []filepermissions.FilePermission {
	return permissions[statkeypatterns.SpecialPermissions.Key]
}

func (fi *FileInfo) getUserPermissions(permissions map[string][]filepermissions.FilePermission) []filepermissions.FilePermission {
	return permissions[statkeypatterns.UserPermissions.Key]
}

func (fi *FileInfo) getGroupPermissions(permissions map[string][]filepermissions.FilePermission) []filepermissions.FilePermission {
	return permissions[statkeypatterns.GroupPermissions.Key]
}

func (fi *FileInfo) getOthersPermissions(permissions map[string][]filepermissions.FilePermission) []filepermissions.FilePermission {
	return permissions[statkeypatterns.OthersPermissions.Key]
}

func (fi *FileInfo) getBirthDateTime() time.Time {
	return parseUnixString(fi.stats[statkeypatterns.BirthDateTime.Key])
}

func (fi *FileInfo) getModificationDateTime() time.Time {
	return parseUnixString(fi.stats[statkeypatterns.ModificationDateTime.Key])
}

func (fi *FileInfo) getChangeDateTime() time.Time {
	return parseUnixString(fi.stats[statkeypatterns.ChangeDateTime.Key])
}

func (fi *FileInfo) getAccessDateTime() time.Time {
	return parseUnixString(fi.stats[statkeypatterns.AccessDateTime.Key])
}

func (fi *FileInfo) getPermissions() string {
	return fi.stats[statkeypatterns.Permissions.Key]
}

func parsePermissions(permissionsStr string) map[string][]filepermissions.FilePermission {
	result := make(map[string][]filepermissions.FilePermission)

	if permissionsStr == "" {
		return result
	}

	chars := []rune(permissionsStr)

	if len(chars) < 4 {
		chars = append([]rune{'0'}, chars...)
	}

	for i := 0; i < 4; i++ {
		perms := make([]filepermissions.FilePermission, 0, 3)

		if chars[i] == '0' {
			continue
		}

		perm := int(chars[i] - '0')

		if i != 0 {
			if perm&4 != 0 {
				perms = append(perms, filepermissions.Read)
			}
			if perm&2 != 0 {
				perms = append(perms, filepermissions.Write)
			}
			if perm&1 != 0 {
				perms = append(perms, filepermissions.Execute)
			}
		} else {
			if perm&4 != 0 {
				perms = append(perms, filepermissions.Setuid)
			}

			if perm&2 != 0 {
				perms = append(perms, filepermissions.Setgid)
			}

			if perm&1 != 0 {
				perms = append(perms, filepermissions.Sticky)
			}
		}

		keys := []string{
			statkeypatterns.SpecialPermissions.Key,
			statkeypatterns.UserPermissions.Key,
			statkeypatterns.GroupPermissions.Key,
			statkeypatterns.OthersPermissions.Key,
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
