// Package archive for simple archive extraction
package archive

import (
	"errors"

	"github.com/markelog/archive/detect"
	"github.com/markelog/archive/tgz"
	"github.com/markelog/archive/zip"
)

// Extract archives
func Extract(src string, dest string) error {
	mime, err := detect.Detect(src)

	if err != nil {
		return err
	}

	switch mime {
	case tgz.Type:
		err = tgz.Extract(src, dest)
	case zip.Type:
		err = zip.Extract(src, dest)
	default:
		err = errors.New("Format is not supported")
	}

	if err != nil {
		return err
	}

	return nil
}
