package interfaces

import (
	"net/http"

	"github.com/sudzekai-web-os/abstractions"
	"github.com/sudzekai-web-os/types"
)

type IFileContentActionsController interface {
	AddRoutes(registry abstractions.IHandlersRegistry)

	// GET files/{filename}/read
	Read(r *http.Request) types.HandlerResult

	// POST files/{filename}/append
	Append(r *http.Request) types.HandlerResult

	// PUT files/{filename}/write
	Write(r *http.Request) types.HandlerResult
}
