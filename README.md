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
            switch err.ErrorType() {
                case cerror.DomainError:
                    log.Fatal(fmt.Sprintf("[%s] %s", err.Code(), err.Error()))
                    break
                case cerror.ApplicationError:
                    log.Fatal(err.ErrorWithTrace())
                    break
                case cerror.BusinessError:
                    log.Fatal(err.Error())
                    break
                }
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

DomainError Log message:
``` shell
[InvalidCastError] Cannot convert the '23x' value to int.
```

ApplicationError Log message:
``` shell
[InvalidCastError] Cannot convert the '23x' value to int.
strconv.Atoi: parsing "23x": invalid syntax
[main.main] /Users/rozturac/go/src/github.com/cerror/_examples/cmd/main.go:32
```

BusinessError Log message:
``` shell
Cannot convert the '23x' value to int.
```

## License

MIT License

Copyright (c) 2021 Rıdvan ÖZTURAÇ

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
