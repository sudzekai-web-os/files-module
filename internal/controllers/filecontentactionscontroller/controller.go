package filecontentactionscontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sudzekai-web-os/core"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/appendfile"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/readfile"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/writefile"
	"github.com/sudzekai-web-os/files-module/internal/errformatter"
	"github.com/sudzekai-web-os/files-module/internal/objects/dto"
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

func (c *Controller) AddRoutes(registry core.IHandlersRegistry) {
	registry.AddHandler("POST /files/write", c.WriteFile)
	registry.AddHandler("POST /files/append", c.AppendFile)
	registry.AddHandler("GET /files/content", c.ReadFile)
}

// POST /files/write?path=...
//
// with body
func (c *Controller) WriteFile(r *http.Request) (result core.HandlerResult) {
	log := c.loggerFactory.NewLogger("controller-POST:files/write")

	path := r.URL.Query().Get("path")

	var body dto.FileContentRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		result.StatusCode = 400
		result.Error = fmt.Errorf("invalid json")
		return
	}

	var cmdResult writefile.CommandResult

	err := mediator.Dispatch(writefile.NewCommand(path, body.Content), &cmdResult)

	if err == nil {
		err = cmdResult.Error
	}
	if err != nil {
		log.LogError("ошибка перезаписи содержимого файла: %s", err.Error())
		return errformatter.MakeBusinessError(err)
	}

	return core.HandlerResult{
		StatusCode: http.StatusOK,
	}
}

// POST /files/append?path=...
//
// with body
func (c *Controller) AppendFile(r *http.Request) (result core.HandlerResult) {
	log := c.loggerFactory.NewLogger("controller-POST:files/append")

	path := r.URL.Query().Get("path")

	var body dto.FileContentRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		result.StatusCode = 400
		result.Error = fmt.Errorf("invalid json")
		return
	}

	var cmdResult appendfile.CommandResult

	err := mediator.Dispatch(appendfile.NewCommand(path, body.Content), &cmdResult)

	if err == nil {
		err = cmdResult.Error
	}
	if err != nil {
		log.LogError("ошибка дополнения содержимого в файл: %s", err.Error())
		return errformatter.MakeBusinessError(err)
	}

	return core.HandlerResult{
		StatusCode: http.StatusOK,
	}
}

// GET /files/content?path=...
func (c *Controller) ReadFile(r *http.Request) (result core.HandlerResult) {
	log := c.loggerFactory.NewLogger("controller-POST:files/append")

	path := r.URL.Query().Get("path")

	var queryResult readfile.QueryResult

	err := mediator.Dispatch(readfile.NewQuery(path), &queryResult)

	if err == nil {
		err = queryResult.Error
	}
	if err != nil {
		log.LogError("ошибка чтения содержимого файла: %s", err.Error())
		return errformatter.MakeBusinessError(err)
	}

	return core.HandlerResult{
		StatusCode: http.StatusOK,
		Data:       queryResult.Content,
	}
}
