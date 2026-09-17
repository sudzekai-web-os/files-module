package filesystem

import (
	"path/filepath"

	"github.com/sudzekai-web-os/abstractions"
	"github.com/sudzekai-web-os/types"
)

type FileSystem struct {
	executor abstractions.IExecutor
}

func New(executor abstractions.IExecutor) *FileSystem {
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

// Deletes file if exists
//
// Errors:
//
//	-
func (fs *FileSystem) Delete(filePath string) error {
	_, err := fs.Stat(filePath)

	if err != nil {
		return err
	}

	cmdResult := fs.executor.Execute("rm", filePath)

	if err := formatCommandResult(cmdResult); err != nil {
		return err
	}

	return nil
}

// Returns bool value representing file existance.
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

// Stat returns information about the specified file.
//
// Errors:
//   - ErrorAbsolutePath
//   - ErrorNotFound
//   - ErrorPermissionDenied
//   - ErrorNotDirectory
//   - ErrorTooManySymlinks
//   - ErrorNameTooLong
//   - ErrorOverflow
//   - ErrorOutOfMemory
//   - ErrorIO
//   - ErrorIsDirectory
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
		return nil, ErrorIsDirectory
	}

	return fi, nil
}

func (fs *FileSystem) IsDir(filePath string) (bool, error) {
	_, err := fs.Stat(filePath)

	if err == nil {
		return false, nil
	}

	if err == ErrorIsDirectory {
		return true, nil
	}

	return false, err
}

func (fs *FileSystem) WriteAllText(filePath, data string) error {
	exists, err := fs.Exists(filePath)

	if err != nil {
		return err
	}

	if !exists {
		return ErrorNotFound
	}

	cmdResult := fs.executor.ExecuteWithInput(data, "tee", filePath)

	if err := formatCommandResult(cmdResult); err != nil {
		return err
	}

	return nil
}

func (fs *FileSystem) AppendAllText(filePath, data string) error {
	exists, err := fs.Exists(filePath)

	if err != nil {
		return err
	}

	if !exists {
		return ErrorNotFound
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

	_, err := fs.Stat(filePath)

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

	cmdResult := fs.executor.Execute("cp", filePath, destinationPath)

	if err := formatCommandResult(cmdResult); err != nil {
		return err
	}

	return nil
}

func formatCommandResult(result types.CommandResult) error {
	if result.Error != nil {
		return ToError(result.Error.Error())
	}

	if result.Stderr != "" {
		return ToError(result.Stderr)
	}

	return nil
}
