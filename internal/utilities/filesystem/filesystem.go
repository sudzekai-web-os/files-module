package filesystem

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/sudzekai-web-os/core"
)

type FileSystem struct {
	executor core.IExecutor
}

func New(executor core.IExecutor) *FileSystem {
	return &FileSystem{
		executor: executor,
	}
}

// Creates file if not exists.
//
// Errors:
//   - ErrorAbsolutePath
//   - ErrorAlreadyExists
//   - Exists func errors
func (fs *FileSystem) Create(filePath string) error {
	exist, err := fs.Exists(filePath)

	if err != nil {
		return err
	}
	if exist {
		return ErrorAlreadyExists
	}

	cmdResult := fs.executor.Execute("touch", filePath)

	if err := formatCommandResult(cmdResult); err != nil {
		return err
	}

	return nil
}

// Deletes file or directory if exists.
func (fs *FileSystem) Delete(filePath string) error {
	fi, err := fs.Stat(filePath)

	if err != nil {
		return err
	}

	var cmdResult core.CommandResult

	if fi.Type == TypeDirectory {
		cmdResult = fs.executor.Execute("rm", "-r", filePath)
	} else {
		cmdResult = fs.executor.Execute("rm", filePath)
	}

	if err := formatCommandResult(cmdResult); err != nil {
		return err
	}

	return nil
}

// Returns bool value representing file or directory existance.
//
// Errors:
//   - ErrorAbsolutePath
//   - Stat func errors
func (fs *FileSystem) Exists(filePath string) (bool, error) {
	if !IsAbsolute(filePath) {
		return false, ErrorAbsolutePath
	}

	result := fs.executor.Execute("test", "-e", filePath)

	if result.Error == nil {
		return true, nil
	}

	if result.Stderr == "" {
		return false, nil
	}

	return false, ToError(result.Stderr)
}

// Stat returns information about the specified file or directory.
//
// Errors:
//   - ErrorAbsolutePath
//   - ErrorNotFound
//   - ErrorPermissionDenied
//   - ErrorTooManySymlinks
//   - ErrorNameTooLong
//   - ErrorOverflow
//   - ErrorOutOfMemory
//   - ErrorIO
//   - Unknown error with provided output
func (fs *FileSystem) Stat(filePath string) (*FileInfo, error) {
	if !IsAbsolute(filePath) {
		return nil, ErrorAbsolutePath
	}

	cmdResult := fs.executor.Execute(
		"stat", "-c", getCommandFullPattern(), filePath,
	)

	if err := formatCommandResult(cmdResult); err != nil {
		return nil, err
	}

	fi := NewFileInfo(cmdResult.Stdout)

	if fi.Type == TypeDirectory {
		fi.FullName = filePath

		cmdResult = fs.executor.Execute("du", "-s", "-B1", filePath)

		if cmdResult.Stdout == "" {
			if err := formatCommandResult(cmdResult); err != nil {
				return nil, err
			}
		}

		value, _ := strconv.ParseInt(strings.Fields(cmdResult.Stdout)[0], 10, 64)

		if value != 0 {
			fi.Size = value
		}
	}

	return fi, nil
}

func (fs *FileSystem) IsDir(filePath string) (bool, error) {
	fi, err := fs.Stat(filePath)

	if err != nil {
		return false, err
	}

	return fi.Type == TypeDirectory, nil
}

func (fs *FileSystem) WriteAllText(filePath, data string) error {
	fi, err := fs.Stat(filePath)

	if err != nil {
		return err
	}

	if fi.Type == TypeDirectory {
		return ErrorIsDirectory
	}

	cmdResult := fs.executor.ExecuteWithInput(data, "tee", filePath)

	if err := formatCommandResult(cmdResult); err != nil {
		return err
	}

	return nil
}

func (fs *FileSystem) AppendAllText(filePath, data string) error {
	fi, err := fs.Stat(filePath)

	if err != nil {
		return err
	}

	if fi.Type == TypeDirectory {
		return ErrorIsDirectory
	}

	cmdResult := fs.executor.ExecuteWithInput(data, "tee", "-a", filePath)

	if err := formatCommandResult(cmdResult); err != nil {
		return err
	}

	return nil
}

func (fs *FileSystem) Rename(filePath, newFileName string) error {
	if !IsAbsolute(filePath) {
		return ErrorAbsolutePath
	}

	return fs.Mv(
		filePath,
		filepath.Join(GetPath(filePath), newFileName),
	)
}

func (fs *FileSystem) Mv(filePath, newFilePath string) error {
	if !IsAbsolute(filePath) || !IsAbsolute(newFilePath) {
		return ErrorAbsolutePath
	}

	_, err := fs.Stat(filePath)

	if err != nil {
		return err
	}

	cmdResult := fs.executor.Execute("mv", filePath, newFilePath)

	if err := formatCommandResult(cmdResult); err != nil {
		return err
	}

	return nil
}

func (fs *FileSystem) Copy(filePath, destinationPath string) error {
	if !IsAbsolute(filePath) || !IsAbsolute(destinationPath) {
		return ErrorAbsolutePath
	}

	fi, err := fs.Stat(filePath)

	if err != nil {
		return err
	}

	exists, err := fs.Exists(destinationPath)

	if err != nil {
		return err
	}

	if exists {
		return ErrorAlreadyExists
	}

	var cmdResult core.CommandResult

	if fi.Type == TypeDirectory {
		cmdResult = fs.executor.Execute("cp", "-r", filePath, destinationPath)
	} else {
		cmdResult = fs.executor.Execute("cp", filePath, destinationPath)
	}

	if err := formatCommandResult(cmdResult); err != nil {
		return err
	}

	return nil
}

func (fs *FileSystem) Read(filePath string) (string, error) {
	fi, err := fs.Stat(filePath)

	if err != nil {
		return "", err
	}

	if fi.Type == TypeDirectory {
		return "", ErrorIsDirectory
	}

	cmdResult := fs.executor.Execute("cat", filePath)

	if err := formatCommandResult(cmdResult); err != nil {
		return "", err
	}

	return cmdResult.Stdout, nil
}

func formatCommandResult(result core.CommandResult) error {
	if result.Stderr != "" {
		return ToError(result.Stderr)
	}

	if result.Error != nil {
		return ToError(result.Error.Error())
	}

	return nil
}
