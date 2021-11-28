#  🥷  CError (Custom Error Handling)

[![GitHub license](https://img.shields.io/github/license/rozturac/cerror.svg?color=24B898&style=for-the-badge&logo=go&logoColor=ffffff)](https://github.com/rozturac/cerror/blob/main/LICENSE)
[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg?color=24B898&style=for-the-badge&logo=go&logoColor=ffffff)](https://pkg.go.dev/github.com/rozturac/cerror)
[![Release](https://img.shields.io/github/tag/rozturac/cerror.svg?label=release&color=24B898&logo=github&style=for-the-badge)](https://github.com/rozturac/cerror/releases/latest)
[![Go Report Card](https://img.shields.io/badge/go%20report-A%2B-green?style=for-the-badge)](https://goreportcard.com/report/github.com/rozturac/cerror)

## Installation

Via go packages:
```go get github.com/rozturac/cerror```

## Usage

Here is a sample CError uses:

```go
import (
    "github.com/rozturac/cerror"
    "log"
    "strconv"
)

func main() {
    defer func() {
        if recover := recover(); recover != nil {
            if err, ok := recover.(cerror.Error); ok {
                log.Fatal(err.ErrorWithTrace())
            }
        }
    }()
    
    x := "23x"
    y, err := strconv.Atoi(x)
    if err != nil {
        panic(cerror.InvalidCastError(x, y).With(err))
    }
}
```

Log message:
``` shell
Cannot convert the '23x' value to a int.
strconv.Atoi: parsing "23x": invalid syntax
[main.main] /Users/rozturac/go/src/github.com/cerror/_examples/cmd/main.go:21
```

Console message:
``` shell
Cannot convert the '23x' value to a int.
```