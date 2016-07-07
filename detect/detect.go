// Package detect detects mimetypes
package detect

import (
	"net/http"
	"os"
)

// Detect mimetype
func Detect(src string) (string, error) {
	file, err := os.Open(src)
	if err != nil {
		return "", err
	}

	defer file.Close()

	data := make([]byte, 512)

	_, err = file.Read(data)

	if err != nil {
		return "", err
	}

	return http.DetectContentType(data), nil
}
