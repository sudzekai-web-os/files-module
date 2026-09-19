package filescontroller

import (
	"net/http"

	"github.com/sudzekai-web-os/core"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/getfileinfo"
	"github.com/sudzekai-web-os/files-module/internal/errformatter"
	"github.com/sudzekai-web-os/mediator"
)

type Controller struct {
	loggerFactory core.ILoggerFactory
}

func New(
	loggerFactory core.ILoggerFactory,
) *Controller {
	return &Controller{
		loggerFactory: loggerFactory,
	}
}

func (fc *Controller) AddRoutes(registry core.IHandlersRegistry) {
	registry.AddHandler("GET /files/info", fc.GetFileInfo)
}

func (fc *Controller) GetFileInfo(r *http.Request) (result core.HandlerResult) {
	log := fc.loggerFactory.NewLogger("controller-GET:/files/info")

	path := r.URL.Query().Get("path")

	var queryResult getfileinfo.QueryResult

	err := mediator.Dispatch(getfileinfo.NewQuery(path), &queryResult)

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
