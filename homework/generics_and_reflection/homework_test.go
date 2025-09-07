package main

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

const (
	PropertyNameTag = "properties"
)

type FieldInfo struct {
	name      string
	omitEmpty bool
}

type Person struct {
	Name    string `properties:"name"`
	Address string `properties:"address,omitempty"`
	Age     int    `properties:"age"`
	Married bool   `properties:"married"`
}

func Serialize[T any](data T) string {
	sb := new(strings.Builder)
	dataType := reflect.TypeOf(data)
	dataValue := reflect.ValueOf(data)
	fieldsCount := dataType.NumField()

	for i := 0; i < fieldsCount; i++ {
		meta := parseMeta(dataType.Field(i))
		if meta == nil {
			continue
		}

		fieldValue := parseFieldValue(dataValue.Field(i))

		if len(fieldValue) > 0 || !meta.omitEmpty {
			sb.WriteString(meta.name)
			sb.WriteString("=")
			sb.WriteString(fieldValue)
			if i < fieldsCount-1 {
				sb.WriteString("\n")
			}
		}
	}

	return sb.String()
}

func parseFieldValue(fieldValue reflect.Value) string {
	switch fieldValue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%v", fieldValue.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%v", fieldValue.Uint())
	case reflect.Float64, reflect.Float32:
		return fmt.Sprintf("%v", fieldValue.Float())
	case reflect.Bool:
		return fmt.Sprintf("%v", fieldValue.Bool())
	default:
		return fieldValue.String()
	}
}

func parseMeta(field reflect.StructField) *FieldInfo {
	props, propExists := field.Tag.Lookup(PropertyNameTag)

	if !propExists || len(props) == 0 {
		return nil
	}

	parts := strings.Split(props, ",")
	omitEmpty := slices.Contains(parts, "omitempty")
	if omitEmpty {
		parts = slices.DeleteFunc(parts, func(s string) bool {
			return s == "omitempty"
		})
	}
	if len(parts) == 0 {
		return nil
	}
	return &FieldInfo{
		name:      parts[0],
		omitEmpty: omitEmpty,
	}
}

func TestSerialization(t *testing.T) {
	tests := map[string]struct {
		person Person
		result string
	}{
		"test case with empty fields": {
			result: "name=\nage=0\nmarried=false",
		},
		"test case with fields": {
			person: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
			},
			result: "name=John Doe\nage=30\nmarried=true",
		},
		"test case with omitempty field": {
			person: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
				Address: "Paris",
			},
			result: "name=John Doe\naddress=Paris\nage=30\nmarried=true",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := Serialize(test.person)
			assert.Equal(t, test.result, result)
		})
	}
}
