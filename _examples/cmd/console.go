package main

import (
	"fmt"
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
					log.Fatal(fmt.Sprintf("[%s] %s", err.ErrorCode(), err.Error()))
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
