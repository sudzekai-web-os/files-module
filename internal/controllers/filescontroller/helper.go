package filescontroller

import (
	"fmt"
	"strings"

	"github.com/sudzekai-web-os/types"
)

func getErroredResult(err error) types.HandlerResult {
	message := err.Error()

	switch {
	case strings.Contains(message, "No such file or directory"):
		return types.HandlerResult{
			StatusCode: 404,
			Error:      fmt.Errorf("такого файла не существует"),
		}

	case strings.Contains(message, "Permission denied"):
		return types.HandlerResult{
			StatusCode: 403,
			Error:      fmt.Errorf("недостаточно прав для доступа к файлу"),
		}

	case strings.Contains(message, "Not a directory"):
		return types.HandlerResult{
			StatusCode: 400,
			Error:      fmt.Errorf("указанный путь содержит некорректный каталог"),
		}

	case strings.Contains(message, "Too many levels of symbolic links"):
		return types.HandlerResult{
			StatusCode: 400,
			Error:      fmt.Errorf("слишком много уровней символических ссылок"),
		}

	case strings.Contains(message, "Value too large for defined data type"):
		return types.HandlerResult{
			StatusCode: 400,
			Error:      fmt.Errorf("путь или файл не поддерживается"),
		}

	default:
		return types.HandlerResult{
			StatusCode: 500,
			Error:      fmt.Errorf("не удалось получить информацию о файле"),
		}
	}
}
