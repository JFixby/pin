package lang

import (
	"fmt"
	"reflect"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func EqualArrays(a, b interface{}) bool {

	arrayA := reflect.ValueOf(a)
	if !(arrayA.Kind() == reflect.Array || arrayA.Kind() == reflect.Slice) {
		ReportErr("Is not array: %v", a)
	}

	arrayB := reflect.ValueOf(b)
	if !(arrayB.Kind() == reflect.Array || arrayB.Kind() == reflect.Slice) {
		ReportErr("Is not array: %v", b)
	}

	La := arrayA.Len()
	Lb := arrayB.Len()
	if La != Lb {
		return false
	}
	return reflect.DeepEqual(a, b)
}

func MinInt(x uint16, y uint16) uint16 {
	if x < y {
		return x
	}
	return y
}

func MaxInt(x uint16, y uint16) uint16 {
	if x < y {
		return y
	}
	return x
}

func ReportErr(format string, a ...interface{}) {
	CheckErr(fmt.Errorf(format, a...))
}

// assertNotNil does sanity check
func AssertNotNil(tag string, value interface{}) {
	if value == nil {
		CheckErr(
			fmt.Errorf("Invalid state: <%v> is nil", tag))
	}
}

func AssertNot(tag string, value interface{}, unexpected interface{}) {
	if value == unexpected {
		CheckErr(
			fmt.Errorf("Invalid state: %v == <%v> ",
				tag,
				unexpected,
			))
	}
}

func AssertValue(tag string, value interface{}, expected interface{}) {
	if value != expected {
		CheckErr(
			fmt.Errorf("Invalid state: %v is <%v>, expected <%v> ",
				tag,
				value,
				expected,
			))
	}
}

// assertNotEmpty does sanity check
func AssertNotEmpty(tag string, value string) {
	if value == "" {
		CheckErr(
			fmt.Errorf("Invalid state: string <%v> is empty", tag))
	}
}
