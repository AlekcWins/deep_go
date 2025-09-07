package main

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type MultiError struct {
	errors []error
}

func (e *MultiError) Error() string {
	if e == nil || len(e.errors) == 0 {
		return ""
	}

	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("%d errors occurred:\n", len(e.errors)))

	for _, err := range e.errors {
		b.WriteString(fmt.Sprintf("\t* %s", err.Error()))
	}

	b.WriteString("\n")
	return b.String()
}

func Append(err error, errs ...error) (res *MultiError) {
	errs = slices.DeleteFunc(errs, func(e error) bool {
		return e == nil
	})

	if err == nil {
		return &MultiError{errors: errs}
	}

	multiError, ok := err.(*MultiError)
	if !ok {
		return &MultiError{
			errors: append([]error{err}, errs...),
		}
	}
	multiError.errors = append(multiError.errors, errs...)

	return multiError
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occurred:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}

func TestMultiError_AppendNilErrors(t *testing.T) {
	err := Append(nil, errors.New("real error"), nil)
	expected := "1 errors occurred:\n\t* real error\n"
	assert.EqualError(t, err, expected)
}

func TestMultiError_AppendToRegularError(t *testing.T) {
	base := errors.New("base error")
	err := Append(base, errors.New("extra 1"), errors.New("extra 2"))

	expected := "3 errors occurred:\n\t* base error\t* extra 1\t* extra 2\n"
	assert.EqualError(t, err, expected)
}
