package detect_test

import (
	"fmt"
	"os"

	"github.com/markelog/archive/tgz"
	"github.com/markelog/archive/zip"

	. "github.com/markelog/archive/detect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Detect", func() {
	pwd, _ := os.Getwd()
	tgzPath := fmt.Sprintf("%v/../testdata/file.tar.gz", pwd)
	zipPath := fmt.Sprintf("%v/../testdata/file.zip", pwd)

	It("detects tar.gz", func() {
		result, _ := Detect(tgzPath)

		Expect(result).To(Equal(tgz.Type))
	})

	It("detects zip", func() {
		result, _ := Detect(zipPath)

		Expect(result).To(Equal(zip.Type))
	})
})
