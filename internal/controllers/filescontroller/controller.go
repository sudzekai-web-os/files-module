package filescontroller

import (
	"net/http"

	"github.com/sudzekai-web-os/abstractions"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/getfileinfo"
	"github.com/sudzekai-web-os/files-module/internal/errformatter"
	"github.com/sudzekai-web-os/mediator"
	"github.com/sudzekai-web-os/types"
)

type Controller struct {
	loggerFactory abstractions.ILoggerFactory
}

func New(
	loggerFactory abstractions.ILoggerFactory,
) *Controller {
	return &Controller{
		loggerFactory: loggerFactory,
	}
}

func (fc *Controller) AddRoutes(registry abstractions.IHandlersRegistry) {
	registry.AddHandler("GET /files/info", fc.GetFileInfo)
}

func (fc *Controller) GetFileInfo(r *http.Request) (result types.HandlerResult) {
	log := fc.loggerFactory.NewLogger("controller-GET:/files/info")

	filePath := r.URL.Query().Get("filePath")

	var queryResult getfileinfo.QueryResult

	err := mediator.Dispatch(getfileinfo.NewQuery(filePath), &queryResult)

	if err == nil {
		err = queryResult.Error
	}
	if err != nil {
		log.LogError("%s", err.Error())
		return errformatter.MakeBusinessError(err)
	}

	result.Data = queryResult.FileInfo
	result.StatusCode = 200

	return
}
