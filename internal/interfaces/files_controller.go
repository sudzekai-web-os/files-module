package interfaces

import (
	"net/http"

	"github.com/sudzekai-web-os/abstractions"
	"github.com/sudzekai-web-os/types"
)

type IFilesController interface {
	AddRoutes(registry abstractions.IHandlersRegistry)

	// GET files/{filename}/fileinfo
	GetFileInfo(r *http.Request) types.HandlerResult
}
