# Archive [![Build Status](https://travis-ci.org/markelog/archive.svg)](https://travis-ci.org/markelog/archive) [![GoDoc](https://godoc.org/github.com/markelog/archive?status.svg)](https://godoc.org/github.com/markelog/archive) [![Go Report Card](https://goreportcard.com/badge/github.com/markelog/archive)](https://goreportcard.com/report/github.com/markelog/archive)

> Simple archive (gzip/zip/bzip2) extraction

## Installation

```
$ go get github.com/markelog/archive
```

## Example

```go
package main

import "github.com/markelog/archive"

func main() {
  // Will extract sexy turtles to current dir
  archive.Extract("/sexy-turtles.tar.gz", ".")

  // Will extract gangsta panda to current dir
  archive.Extract("/gangsta-panda.zip", ".")

  // Will extract killa gorilla to current dir
  archive.Extract("/killa-gorilla.tar.bz2", ".")
}
```
