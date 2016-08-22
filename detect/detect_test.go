package detect_test

import (
	"fmt"
	"os"

	"github.com/markelog/archive/tgz"
	"github.com/markelog/archive/zip"
	"github.com/markelog/archive/bz2"

	. "github.com/markelog/archive/detect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Detect", func() {
	pwd, _ := os.Getwd()
	tgzPath := fmt.Sprintf("%v/../testdata/file.tar.gz", pwd)
	zipPath := fmt.Sprintf("%v/../testdata/file.zip", pwd)
	bz2Path := fmt.Sprintf("%v/../testdata/file.tar.bz2", pwd)

	It("detects tar.bz2", func() {
		result, _ := Detect(bz2Path)

		Expect(result).To(Equal(bz2.Type))
	})

	It("detects tar.gz", func() {
		result, _ := Detect(tgzPath)

		Expect(result).To(Equal(tgz.Type))
	})

	It("detects zip", func() {
		result, _ := Detect(zipPath)

		Expect(result).To(Equal(zip.Type))
	})
})
