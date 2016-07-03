package archive_test

import (
	"fmt"
	"io/ioutil"
	"os"

	"path/filepath"

	. "github.com/markelog/archive"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func isExist(src string) bool {
	_, err := os.Stat(src)
	return err == nil
}

var _ = Describe("Archive", func() {
	pwd, _ := os.Getwd()
	file := fmt.Sprintf("%v/testdata/file.tar.gz", pwd)
	folder := fmt.Sprintf("%v/testdata/folder.tar.gz", pwd)
	var (
		tmpFolder string
	)

	Describe("Extract tar.gz", func() {
		BeforeEach(func() {
			tmpFolder, _ = ioutil.TempDir("", "tmp")
		})

		AfterEach(func() {
			os.Remove(tmpFolder)
		})

		Describe("file", func() {
			var (
				path string
			)

			BeforeEach(func() {
				Extract(file, tmpFolder)
				path = filepath.Join(tmpFolder, "1")
			})

			It("should exist", func() {
				Expect(isExist(path)).To(Equal(true))
			})

			It("should get file content", func() {
				data, _ := ioutil.ReadFile(path)

				Expect(string(data)).To(Equal("test\n"))
			})
		})

		Describe("folder", func() {
			var (
				path string
			)

			BeforeEach(func() {
				Extract(folder, tmpFolder)
				path = filepath.Join(tmpFolder, "test")
			})

			It("should exist", func() {
				Expect(isExist(path)).To(Equal(true))
			})

			It("should get folder content", func() {
				data, _ := ioutil.ReadFile(path + "/1")

				Expect(string(data)).To(Equal("test\n"))
			})
		})
	})
})
