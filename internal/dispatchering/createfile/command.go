package createfile

type Command struct {
	FilePath string
}

func NewCommand(filePath string) Command {
	return Command{
		FilePath: filePath,
	}
}
