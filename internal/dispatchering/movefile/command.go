package movefile

type Command struct {
	FilePath        string
	DestinationPath string
}

func NewCommand(filePath, destinationPath string) Command {
	return Command{
		FilePath:        filePath,
		DestinationPath: destinationPath,
	}
}
