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

  tgzFile := fmt.Sprintf("%v/testdata/file.tar.gz", pwd)
  tgzFolder := fmt.Sprintf("%v/testdata/folder.tar.gz", pwd)

  zipFile := fmt.Sprintf("%v/testdata/file.zip", pwd)
  zipFolder := fmt.Sprintf("%v/testdata/folder.zip", pwd)

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
        Extract(tgzFile, tmpFolder)
        path = filepath.Join(tmpFolder, "file")
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
        Extract(tgzFolder, tmpFolder)
        path = filepath.Join(tmpFolder, "folder")
      })

      It("should exist", func() {
        Expect(isExist(path)).To(Equal(true))
      })

      It("should get folder content", func() {
        data, _ := ioutil.ReadFile(path + "/file")

        Expect(string(data)).To(Equal("test\n"))
      })
    })
  })

  Describe("Extract zip", func() {
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
        Extract(zipFile, tmpFolder)
        path = filepath.Join(tmpFolder, "file")
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
        Extract(zipFolder, tmpFolder)
        path = filepath.Join(tmpFolder, "folder")
      })

      It("should exist", func() {
        Expect(isExist(path)).To(Equal(true))
      })

      It("should get folder content", func() {
        data, _ := ioutil.ReadFile(path + "/file")

        Expect(string(data)).To(Equal("test\n"))
      })
    })
  })
})
