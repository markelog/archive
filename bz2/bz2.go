// Package bz2 extracts bzip2
package bz2

import (
	"archive/tar"
	"compress/bzip2"
	"path/filepath"

	"io"
	"os"
)

// Mimetype of the tarball
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

	reader := bzip2.NewReader(file)
	tarReader := tar.NewReader(reader)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if header.Name == "." {
			continue
		}

		err = extract(header, dest, tarReader)
		if err != nil {
			return err
		}
	}

	return nil
}

func extract(header *tar.Header, dest string, reader io.Reader) error {
	path := filepath.Join(dest, header.Name)
	info := header.FileInfo()

	if info.IsDir() {
		err := os.MkdirAll(path, info.Mode())
		if err != nil {
			return err
		}

		return nil
	}

	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	if info.Mode()&os.ModeSymlink != 0 {
		return os.Symlink(header.Linkname, path)
	}

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, info.Mode())

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, reader)

	if err != nil {
		return err
	}

	return nil
}
