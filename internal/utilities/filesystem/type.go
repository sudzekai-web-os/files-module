package filesystem

type FileType string

const (
	TypeRegular         FileType = "file"
	TypeDirectory       FileType = "directory"
	TypeSymlink         FileType = "symlink"
	TypeSocket          FileType = "socket"
	TypePipe            FileType = "pipe"
	TypeBlockDevice     FileType = "block device"
	TypeCharacterDevice FileType = "character device"
	TypeUnknown         FileType = "unknown"
)
