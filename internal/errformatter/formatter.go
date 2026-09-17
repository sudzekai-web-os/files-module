package errformatter

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sudzekai-web-os/types"
)

var errors = map[string]types.HandlerResult{
	"No such file or directory": {
		StatusCode: http.StatusNotFound,
		Error:      fmt.Errorf("такого файла не существует"),
	},
	"Permission denied": {
		StatusCode: http.StatusForbidden,
		Error:      fmt.Errorf("недостаточно прав для доступа к файлу"),
	},
	"Not a directory": {
		StatusCode: http.StatusBadRequest,
		Error:      fmt.Errorf("указанный путь содержит некорректный каталог"),
	},
	"Too many levels of symbolic links": {
		StatusCode: http.StatusBadRequest,
		Error:      fmt.Errorf("слишком много уровней символических ссылок"),
	},
	"File name too long": {
		StatusCode: http.StatusBadRequest,
		Error:      fmt.Errorf("имя файла слишком длинное"),
	},
	"Value too large for defined data type": {
		StatusCode: http.StatusBadRequest,
		Error:      fmt.Errorf("путь или файл не поддерживается"),
	},
	"Cannot allocate memory": {
		StatusCode: http.StatusInternalServerError,
		Error:      fmt.Errorf("не удалось выделить память для выполнения операции"),
	},
	"Input/output error": {
		StatusCode: http.StatusInternalServerError,
		Error:      fmt.Errorf("ошибка ввода-вывода"),
	},
	"Is a directory": {
		StatusCode: http.StatusBadRequest,
		Error:      fmt.Errorf("указанный путь является каталогом"),
	},
	"Provided path is pointer to directory": {
		StatusCode: http.StatusBadRequest,
		Error:      fmt.Errorf("указанный путь является каталогом"),
	},
	"Read-only file system": {
		StatusCode: http.StatusForbidden,
		Error:      fmt.Errorf("файловая система доступна только для чтения"),
	},
	"File system is read-only": {
		StatusCode: http.StatusForbidden,
		Error:      fmt.Errorf("файловая система доступна только для чтения"),
	},
	"Device or resource busy": {
		StatusCode: http.StatusConflict,
		Error:      fmt.Errorf("файл или устройство занято"),
	},
	"No space left on device": {
		StatusCode: http.StatusInsufficientStorage,
		Error:      fmt.Errorf("недостаточно свободного места"),
	},
	"Disk quota exceeded": {
		StatusCode: http.StatusInsufficientStorage,
		Error:      fmt.Errorf("превышена квота дискового пространства"),
	},
	"Operation not permitted": {
		StatusCode: http.StatusForbidden,
		Error:      fmt.Errorf("операция запрещена"),
	},
	"Invalid argument": {
		StatusCode: http.StatusBadRequest,
		Error:      fmt.Errorf("передан некорректный аргумент"),
	},
	"Interrupted system call": {
		StatusCode: http.StatusInternalServerError,
		Error:      fmt.Errorf("операция была прервана"),
	},
	"Bad file descriptor": {
		StatusCode: http.StatusInternalServerError,
		Error:      fmt.Errorf("некорректный файловый дескриптор"),
	},
	"Too many open files": {
		StatusCode: http.StatusInternalServerError,
		Error:      fmt.Errorf("превышено количество открытых файлов"),
	},
	"File too large": {
		StatusCode: http.StatusRequestEntityTooLarge,
		Error:      fmt.Errorf("файл слишком большой"),
	},
	"Invalid cross-device link": {
		StatusCode: http.StatusConflict,
		Error:      fmt.Errorf("операция между файловыми системами не поддерживается"),
	},
	"Directory not empty": {
		StatusCode: http.StatusConflict,
		Error:      fmt.Errorf("каталог не пуст"),
	},
	"File already exists": {
		StatusCode: http.StatusConflict,
		Error:      fmt.Errorf("файл уже существует"),
	},
	"Path must be absolute": {
		StatusCode: http.StatusBadRequest,
		Error:      fmt.Errorf("путь должен быть абсолютным"),
	},
}

var defaultErr = types.HandlerResult{
	StatusCode: http.StatusInternalServerError,
	Error:      fmt.Errorf("неизвестная ошибка сервера"),
}

func MakeBusinessError(err error) types.HandlerResult {
	errString := err.Error()

	for sysErr, busErr := range errors {
		if strings.Contains(errString, sysErr) {
			return busErr
		}
	}

	return defaultErr
}
