package filesystem

import (
	"fmt"
	"strings"
)

var (
	ErrorAlreadyExists         = fmt.Errorf("File already exists")
	ErrorAbsolutePath          = fmt.Errorf("Path must be absolute")
	ErrorNotFound              = fmt.Errorf("File not found")
	ErrorPermissionDenied      = fmt.Errorf("Permission denied")
	ErrorNotDirectory          = fmt.Errorf("Not a directory")
	ErrorTooManySymlinks       = fmt.Errorf("Too many levels of symbolic links")
	ErrorNameTooLong           = fmt.Errorf("File name too long")
	ErrorOverflow              = fmt.Errorf("Value too large for defined data type")
	ErrorOutOfMemory           = fmt.Errorf("Cannot allocate memory")
	ErrorIO                    = fmt.Errorf("Input/output error")
	ErrorIsDirectory           = fmt.Errorf("Provided path is pointer to directory")
	ErrorReadOnlyFileSystem    = fmt.Errorf("Read-only file system")
	ErrorResourceBusy          = fmt.Errorf("Device or resource busy")
	ErrorNoSpace               = fmt.Errorf("No space left on device")
	ErrorQuotaExceeded         = fmt.Errorf("Disk quota exceeded")
	ErrorOperationNotPermitted = fmt.Errorf("Operation not permitted")
	ErrorInvalidArgument       = fmt.Errorf("Invalid argument")
	ErrorInterrupted           = fmt.Errorf("Interrupted system call")
	ErrorBadFileDescriptor     = fmt.Errorf("Bad file descriptor")
	ErrorTooManyOpenFiles      = fmt.Errorf("Too many open files")
	ErrorFileTooLarge          = fmt.Errorf("File too large")
	ErrorCrossDevice           = fmt.Errorf("Invalid cross-device link")
	ErrorDirectoryNotEmpty     = fmt.Errorf("Directory not empty")
)

var errors = map[string]error{
	"No such file or directory":             ErrorNotFound,
	"Permission denied":                     ErrorPermissionDenied,
	"Not a directory":                       ErrorNotDirectory,
	"Too many levels of symbolic links":     ErrorTooManySymlinks,
	"File name too long":                    ErrorNameTooLong,
	"Value too large for defined data type": ErrorOverflow,
	"Cannot allocate memory":                ErrorOutOfMemory,
	"Input/output error":                    ErrorIO,
	"Read-only file system":                 ErrorReadOnlyFileSystem,
	"File system is read-only":              ErrorReadOnlyFileSystem,
	"Is a directory":                        ErrorIsDirectory,
	"Device or resource busy":               ErrorResourceBusy,
	"No space left on device":               ErrorNoSpace,
	"Disk quota exceeded":                   ErrorQuotaExceeded,
	"Operation not permitted":               ErrorOperationNotPermitted,
	"Invalid argument":                      ErrorInvalidArgument,
	"Interrupted system call":               ErrorInterrupted,
	"Bad file descriptor":                   ErrorBadFileDescriptor,
	"Too many open files":                   ErrorTooManyOpenFiles,
	"File too large":                        ErrorFileTooLarge,
	"Invalid cross-device link":             ErrorCrossDevice,
	"Directory not empty":                   ErrorDirectoryNotEmpty,
}

func ToError(stdErr string) error {
	for stdErrSubStr, err := range errors {
		if strings.Contains(stdErr, stdErrSubStr) {
			return err
		}
	}

	return fmt.Errorf("%s", stdErr)
}
