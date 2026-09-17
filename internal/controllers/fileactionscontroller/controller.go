package fileactionscontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sudzekai-web-os/abstractions"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/copyfile"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/createfile"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/deletefile"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/movefile"
	"github.com/sudzekai-web-os/files-module/internal/dispatchering/renamefile"
	"github.com/sudzekai-web-os/files-module/internal/errformatter"
	"github.com/sudzekai-web-os/files-module/internal/objects/dto"
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

func (fac *Controller) AddRoutes(registry abstractions.IHandlersRegistry) {
	registry.AddHandler("POST /files", fac.CreateFile)
	registry.AddHandler("DELETE /files", fac.DeleteFile)
	registry.AddHandler("PATCH /files/name", fac.RenameFile)
	registry.AddHandler("PATCH /files/location", fac.MoveFile)
	registry.AddHandler("POST /files/copy", fac.CopyFile)
}

// POST /files?path=..
func (fac *Controller) CreateFile(r *http.Request) (result types.HandlerResult) {
	log := fac.loggerFactory.NewLogger("controller-POST:files/")

	path := r.URL.Query().Get("path")

	var cmdResult createfile.CommandResult

	err := mediator.Dispatch(createfile.NewCommand(path), &cmdResult)

	if err == nil {
		err = cmdResult.Error
	}
	if err != nil {
		log.LogError("ошибка создания файла: %s", err.Error())
		return errformatter.MakeBusinessError(err)
	}

	return types.HandlerResult{
		StatusCode: http.StatusCreated,
	}
}

// DELETE /files?path=..
func (fac *Controller) DeleteFile(r *http.Request) (result types.HandlerResult) {
	log := fac.loggerFactory.NewLogger("controller-DELETE:files/")

	path := r.URL.Query().Get("path")

	var cmdResult deletefile.CommandResult

	err := mediator.Dispatch(deletefile.NewCommand(path), &cmdResult)

	if err == nil {
		err = cmdResult.Error
	}
	if err != nil {
		log.LogError("ошибка удаления файла: %s", err.Error())
		return errformatter.MakeBusinessError(err)
	}

	return types.HandlerResult{
		StatusCode: http.StatusOK,
	}
}

// PATCH /files/name?path=..
//
// with body
func (fac *Controller) RenameFile(r *http.Request) (result types.HandlerResult) {
	log := fac.loggerFactory.NewLogger("controller-PATCH:files/name")

	path := r.URL.Query().Get("path")

	var body dto.FileRenameRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		result.StatusCode = 400
		result.Error = fmt.Errorf("invalid json")
		return
	}

	var cmdResult renamefile.CommandResult

	err := mediator.Dispatch(renamefile.NewCommand(path, body.NewFileName), &cmdResult)

	if err == nil {
		err = cmdResult.Error
	}
	if err != nil {
		log.LogError("ошибка переименования файла: %s", err.Error())
		return errformatter.MakeBusinessError(err)
	}

	return types.HandlerResult{
		StatusCode: http.StatusOK,
	}
}

// PATCH /files/location?path=..
//
// with body
func (fac *Controller) MoveFile(r *http.Request) (result types.HandlerResult) {
	log := fac.loggerFactory.NewLogger("controller-PATCH:files/location")

	path := r.URL.Query().Get("path")

	var body dto.FileMoveRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		result.StatusCode = 400
		result.Error = fmt.Errorf("invalid json")
		return
	}

	var cmdResult movefile.CommandResult

	err := mediator.Dispatch(movefile.NewCommand(path, body.DestinationPath), &cmdResult)

	if err == nil {
		err = cmdResult.Error
	}
	if err != nil {
		log.LogError("ошибка удаления файла: %s", err.Error())
		return errformatter.MakeBusinessError(err)
	}

	return types.HandlerResult{
		StatusCode: http.StatusOK,
	}
}

// POST /files/copy?path=..
//
// with body
func (fac *Controller) CopyFile(r *http.Request) (result types.HandlerResult) {
	log := fac.loggerFactory.NewLogger("controller-POST:files/copy")

	path := r.URL.Query().Get("path")

	var body dto.FileMoveRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		result.StatusCode = 400
		result.Error = fmt.Errorf("invalid json")
		return
	}

	var cmdResult copyfile.CommandResult

	err := mediator.Dispatch(copyfile.NewCommand(path, body.DestinationPath), &cmdResult)

	if err == nil {
		err = cmdResult.Error
	}
	if err != nil {
		log.LogError("ошибка копирования файла: %s", err.Error())
		return errformatter.MakeBusinessError(err)
	}

	return types.HandlerResult{
		StatusCode: http.StatusCreated,
	}
}
