package interfaces

import (
	"net/http"

	"github.com/sudzekai-web-os/abstractions"
	"github.com/sudzekai-web-os/types"
)

type IFileActionsController interface {
	AddRoutes(srv abstractions.IServer)

	// POST files/
	CreateFile(r *http.Request) types.HandlerResult

	// DELETE files/{filename}
	DeleteFile(r *http.Request) types.HandlerResult

	// PATCH files/{filename}
	ChangeFileName(r *http.Request) types.HandlerResult
}
