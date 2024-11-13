package httpUtils

import (
	"strings"
)

var imgExtensions = map[string]bool{
	"jpeg": true,
	"jpg":  true,
	"png":  true,
}

func GetFileExtension(fileName string) string {
	splitted := strings.Split(fileName, ".")
	extension := splitted[len(splitted)-1]

	return extension
}

func IsImageExtensionAllowed(fileName string) bool {
	extension := GetFileExtension(fileName)

	ext := imgExtensions[extension]

	return ext
}
