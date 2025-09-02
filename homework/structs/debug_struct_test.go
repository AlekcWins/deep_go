package structs

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

// Получение имени поля для заданного смещения
func getFieldName(typ reflect.Type, offset uintptr) string {
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		fieldStart := f.Offset
		fieldEnd := fieldStart + f.Type.Size()
		if offset >= fieldStart && offset < fieldEnd {
			// Если это структура, рекурсивно ищем вложенные поля
			if f.Type.Kind() == reflect.Struct {
				name := getFieldName(f.Type, offset-fieldStart)
				if name != "" {
					return f.Name + "." + name
				}
			}
			return f.Name
		}
	}
	return "padding"
}

// Печать битов структуры с указанием поля и выравнивания
func dumpStructBits(v interface{}) {
	val := reflect.ValueOf(v).Elem() // нужно Elem() для указателя
	typ := val.Type()
	size := typ.Size()
	ptr := unsafe.Pointer(val.UnsafeAddr())
	data := (*[1 << 16]byte)(ptr)

	fmt.Printf("Struct %s, size: %d bytes\n", typ.Name(), size)
	fmt.Printf("Byte | Bits       | Field\n")
	fmt.Printf("-----+------------+----------------\n")

	for i := uintptr(0); i < size; i++ {
		fieldName := getFieldName(typ, i)
		b := data[i]
		fmt.Printf("%3d  | ", i)
		for j := 7; j >= 0; j-- {
			bit := (b >> j) & 1
			fmt.Printf("%d", bit)
		}
		fmt.Printf(" | %s\n", fieldName)
	}
}

func TestGamePersonBitMapReflection(t *testing.T) {
	var gp GamePerson
	dumpStructBits(&gp)
}
