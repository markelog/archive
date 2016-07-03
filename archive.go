// Simple tarball extraction
package archive

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

func Extract(src, dest string) error {
	open, err := os.Open(src)

	if err != nil {
		return err
	}

	defer open.Close()

	reader, _ := gzip.NewReader(open)
	defer reader.Close()

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

		err = extractTar(header, dest, tarReader)
		if err != nil {
			return err
		}
	}

	return nil
}

func extractTar(header *tar.Header, dest string, input io.Reader) error {
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

	_, err = io.Copy(file, input)

	if err != nil {
		return err
	}

	return nil
}
