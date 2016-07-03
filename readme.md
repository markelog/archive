[![Build Status](https://travis-ci.org/markelog/archive.svg?branch=master)](https://travis-ci.org/markelog/archive)

# Archive

Simple tarball extraction.

Check the [docs](http://godoc.org/github.com/markelog/archive)

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
}
```
