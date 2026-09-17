package movefile

type CommandResult struct {
	Error error
}

func NewResult(err error) CommandResult {
	return CommandResult{
		Error: err,
	}
}
