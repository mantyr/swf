# SWF Adobe Flash file (only info about file)

[![Build Status](https://travis-ci.org/mantyr/swf.svg?branch=master)](https://travis-ci.org/mantyr/swf) [![GoDoc](https://godoc.org/github.com/mantyr/swf?status.png)](http://godoc.org/github.com/mantyr/swf) [![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE.md)

## Installation

    $ go get github.com/mantyr/swf

## Example

```Go
package main

import (
    "github.com/mantyr/swf"
    "fmt"
)

func main() {
    s, err := swf.Open("./testdata/f0673319418b054b4f108c18736a042f5dcec4e2123401c48a727dfaab7354ef.swf")
    fmt.Println(s.Width())  // 640
    fmt.Println(s.Height()) // 480
}
```

## Author

[Oleg Shevelev][mantyr]

[mantyr]: https://github.com/mantyr