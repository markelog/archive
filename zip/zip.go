// Simple zip extraction
package zip

import (
  "archive/zip"
  "io"
  "io/ioutil"
  "os"
  "path/filepath"
)

var (
  Type = "application/zip"
)

func Extract(src string, dest string) error {
  files, err := zip.OpenReader(src)

  if err != nil {
    return err
  }

  defer files.Close()

  for _, file := range files.File {
    err = extractZip(file, dest)

    if err != nil {
      return err
    }
  }

  return nil
}

func extractZip(file *zip.File, dest string) error {
  reader, err := file.Open()

  if err != nil {
    return err
  }

  defer reader.Close()

  return extract(file, dest, reader)
}

func extract(header *zip.File, dest string, reader io.Reader) error {
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
    linkName, err := ioutil.ReadAll(reader)

    if err != nil {
      return err
    }

    return os.Symlink(string(linkName), path)
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
