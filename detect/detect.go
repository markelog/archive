// Package detect detects mimetypes
package detect

import (
	"strings"
	"net/http"
	"os"
)

func isbz2(src string) bool {
	return strings.HasSuffix(src, "bz2")
}

// Detect mimetype
func Detect(src string) (string, error) {
	// Apparently, golang can't detect bzip mimetype :/

	if isbz2(src) {
		return "application/x-bzip2", nil
	}

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
