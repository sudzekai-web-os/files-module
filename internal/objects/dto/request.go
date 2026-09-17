package dto

type FileRenameRequest struct {
	NewFileName string
}

type FileMoveRequest struct {
	DestinationPath string
}

type FileContentRequest struct {
	Content string
}
