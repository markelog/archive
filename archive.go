// Package archive for simple archive extraction
package archive

import (
	"errors"

	"github.com/markelog/archive/detect"

	"github.com/markelog/archive/gzip"
	"github.com/markelog/archive/zip"
	"github.com/markelog/archive/bz2"
)

// Extract archives
func Extract(src string, dest string) error {
	mime, err := detect.Detect(src)

	if err != nil {
		return err
	}

	switch mime {
	case gzip.Type:
		err = gzip.Extract(src, dest)
	case zip.Type:
		err = zip.Extract(src, dest)
	case bz2.Type:
		err = bz2.Extract(src, dest)
	default:
		err = errors.New("Format is not supported")
	}

	if err != nil {
		return err
	}

	return nil
}
