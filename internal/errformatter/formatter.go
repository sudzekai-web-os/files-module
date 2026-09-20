package errformatter

import (
	"fmt"
	"net/http"

	"github.com/sudzekai-web-os/core"
	"github.com/sudzekai-web-os/files-module/internal/utilities/filesystem"
)

var errors = map[error]core.HandlerResult{
	filesystem.ErrorAlreadyExists: {
		StatusCode: http.StatusConflict,
		Error:      fmt.Errorf("файл уже существует"),
	},
	filesystem.ErrorAbsolutePath: {
		StatusCode: http.StatusBadRequest,
		Error:      fmt.Errorf("путь должен быть абсолютным"),
	},
	filesystem.ErrorNotFound: {
		StatusCode: http.StatusNotFound,
		Error:      fmt.Errorf("указанный файл не существует"),
	},
	filesystem.ErrorPermissionDenied: {
		StatusCode: http.StatusForbidden,
		Error:      fmt.Errorf("недостаточно прав для доступа к файлу"),
	},
	filesystem.ErrorNotDirectory: {
		StatusCode: http.StatusBadRequest,
		Error:      fmt.Errorf("указанный путь не является каталогом"),
	},
	filesystem.ErrorTooManySymlinks: {
		StatusCode: http.StatusBadRequest,
		Error:      fmt.Errorf("слишком много уровней символических ссылок"),
	},
	filesystem.ErrorNameTooLong: {
		StatusCode: http.StatusBadRequest,
		Error:      fmt.Errorf("имя файла слишком длинное"),
	},
	filesystem.ErrorOverflow: {
		StatusCode: http.StatusBadRequest,
		Error:      fmt.Errorf("путь или файл не поддерживается"),
	},
	filesystem.ErrorOutOfMemory: {
		StatusCode: http.StatusInternalServerError,
		Error:      fmt.Errorf("не удалось выделить память"),
	},
	filesystem.ErrorIO: {
		StatusCode: http.StatusInternalServerError,
		Error:      fmt.Errorf("ошибка ввода-вывода"),
	},
	filesystem.ErrorIsDirectory: {
		StatusCode: http.StatusBadRequest,
		Error:      fmt.Errorf("указанный путь является каталогом"),
	},
	filesystem.ErrorReadOnlyFileSystem: {
		StatusCode: http.StatusForbidden,
		Error:      fmt.Errorf("файловая система доступна только для чтения"),
	},
	filesystem.ErrorResourceBusy: {
		StatusCode: http.StatusConflict,
		Error:      fmt.Errorf("файл или устройство занято"),
	},
	filesystem.ErrorNoSpace: {
		StatusCode: http.StatusInsufficientStorage,
		Error:      fmt.Errorf("недостаточно свободного места"),
	},
	filesystem.ErrorQuotaExceeded: {
		StatusCode: http.StatusInsufficientStorage,
		Error:      fmt.Errorf("превышена квота дискового пространства"),
	},
	filesystem.ErrorOperationNotPermitted: {
		StatusCode: http.StatusForbidden,
		Error:      fmt.Errorf("операция запрещена"),
	},
	filesystem.ErrorInvalidArgument: {
		StatusCode: http.StatusBadRequest,
		Error:      fmt.Errorf("передан некорректный аргумент"),
	},
	filesystem.ErrorInterrupted: {
		StatusCode: http.StatusInternalServerError,
		Error:      fmt.Errorf("операция была прервана"),
	},
	filesystem.ErrorBadFileDescriptor: {
		StatusCode: http.StatusInternalServerError,
		Error:      fmt.Errorf("некорректный файловый дескриптор"),
	},
	filesystem.ErrorTooManyOpenFiles: {
		StatusCode: http.StatusInternalServerError,
		Error:      fmt.Errorf("превышено количество открытых файлов"),
	},
	filesystem.ErrorFileTooLarge: {
		StatusCode: http.StatusRequestEntityTooLarge,
		Error:      fmt.Errorf("файл слишком большой"),
	},
	filesystem.ErrorCrossDevice: {
		StatusCode: http.StatusConflict,
		Error:      fmt.Errorf("операция между файловыми системами не поддерживается"),
	},
	filesystem.ErrorDirectoryNotEmpty: {
		StatusCode: http.StatusConflict,
		Error:      fmt.Errorf("каталог не пуст"),
	},
	filesystem.ErrorTooManyLinks: {
		StatusCode: http.StatusConflict,
		Error:      fmt.Errorf("превышено количество ссылок"),
	},
	filesystem.ErrorNoSuchDevice: {
		StatusCode: http.StatusNotFound,
		Error:      fmt.Errorf("устройство не найдено"),
	},
	filesystem.ErrorNoSuchDeviceOrAddress: {
		StatusCode: http.StatusNotFound,
		Error:      fmt.Errorf("устройство или адрес не найден"),
	},
	filesystem.ErrorStaleFileHandle: {
		StatusCode: http.StatusInternalServerError,
		Error:      fmt.Errorf("файловый дескриптор устарел"),
	},
}

var defaultErr = core.HandlerResult{
	StatusCode: http.StatusInternalServerError,
	Error:      fmt.Errorf("неизвестная ошибка сервера"),
}

func MakeBusinessError(err error) core.HandlerResult {
	if result, exists := errors[err]; exists {
		return result
	}

	return defaultErr
}
