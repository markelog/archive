// Package tar extracts tarballs
package tar

import (
	"archive/tar"
	"path/filepath"

	"io"
	"os"
)

// Extract tarballs
func Extract(reader io.Reader, dest string) error {
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

		err = extract(tarReader, header, dest)
		if err != nil {
			return err
		}
	}

	return nil
}

func extract(reader io.Reader, header *tar.Header, dest string) error {
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
