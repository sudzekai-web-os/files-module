package renamefile

type Command struct {
	FilePath    string
	NewFileName string
}

func NewCommand(filePath, newFileName string) Command {
	return Command{
		FilePath:    filePath,
		NewFileName: newFileName,
	}
}
