package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go
type MyInt interface {
	~uint32 | ~uint16 | ~uint64
}

func ToLittleEndian[T MyInt](number T) T {
	startP := unsafe.Pointer(&number)
	size := int(unsafe.Sizeof(number))
	endP := unsafe.Add(startP, size-1)
	for i := 0; i < size/2; i++ {
		firsByte := (*byte)(unsafe.Add(startP, i))
		lastByte := (*byte)(unsafe.Add(endP, -i))
		*firsByte, *lastByte = *lastByte, *firsByte
	}
	return number
}

func TestConversion(t *testing.T) {
	tests := map[string]struct {
		number uint32
		result uint32
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x04030201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}

	uint16Tests := map[string]struct {
		number uint16
		result uint16
	}{
		"uint16 #1": {number: 0x0000, result: 0x0000},
		"uint16 #2": {number: 0xFFFF, result: 0xFFFF},
		"uint16 #3": {number: 0x1234, result: 0x3412},
		"uint16 #4": {number: 0x00FF, result: 0xFF00},
	}

	for name, test := range uint16Tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}

	uint64Tests := map[string]struct {
		number uint64
		result uint64
	}{
		"uint64 #1": {number: 0x0000000000000000, result: 0x0000000000000000},
		"uint64 #2": {number: 0xFFFFFFFFFFFFFFFF, result: 0xFFFFFFFFFFFFFFFF},
		"uint64 #3": {number: 0x0102030405060708, result: 0x0807060504030201},
		"uint64 #4": {number: 0x00000000FFFFFFFF, result: 0xFFFFFFFF00000000},
	}

	for name, test := range uint64Tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
