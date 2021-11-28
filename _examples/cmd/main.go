package main

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
