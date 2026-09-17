package appendfile

type Command struct {
	FilePath string
	Data     string
}

func NewCommand(filePath, data string) Command {
	return Command{
		FilePath: filePath,
		Data:     data,
	}
}
