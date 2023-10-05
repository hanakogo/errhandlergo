package test

import (
	"fmt"
	"github.com/ohanakogo/errhandlergo"
	"testing"
)

func TestErrorHandler(t *testing.T) {
	err := fmt.Errorf("test err")

	handler := errhandlergo.NewDefaultErrorHandler()
	handler.H(err)

	handler.SetHandler(func(err error) {
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("custom error handler")
		}
	})
	handler.H(err)
}
