package controllerhelper

import (
	"fmt"
	"path/filepath"
	"strings"
)

func ValidateFilePath(path string) error {
	if path == "" {
		return fmt.Errorf("путь к файлу не может быть пустым")
	}

	if !filepath.IsAbs(path) {
		return fmt.Errorf("путь к файлу должен быть абсолютным")
	}

	if strings.HasSuffix(path, "/") {
		return fmt.Errorf("путь к файлу не может заканчиваться на '/'")
	}

	if strings.ContainsRune(path, '\x00') {
		return fmt.Errorf("путь к файлу содержит недопустимый символ")
	}

	if path == "/" {
		return fmt.Errorf("путь должен указывать на файл, а не на корневой каталог")
	}

	return nil
}
