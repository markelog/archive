// Package gzip extracts gzips
package gzip

import (
	"compress/gzip"
	"os"

	"github.com/markelog/archive/tar"
)

// Mimetype of the gzip
var (
	Type = "application/x-gzip"
)

// Extract gzip
func Extract(src string, dest string) error {
	file, err := os.Open(src)

	if err != nil {
		return err
	}

	defer file.Close()

	reader, _ := gzip.NewReader(file)
	defer reader.Close()

	return tar.Extract(reader, dest)
}
