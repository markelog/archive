// Package bz2 extracts bzip2
package bz2

import (
	"compress/bzip2"
	"os"

	"github.com/markelog/archive/tar"
)

// Mimetype of the bzip2
var (
	Type = "application/x-bzip2"
)

// Extract bzip2
func Extract(src string, dest string) error {
	file, err := os.Open(src)

	if err != nil {
		return err
	}

	defer file.Close()

	return tar.Extract(bzip2.NewReader(file), dest)
}
