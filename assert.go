package assert

import (
	"reflect"
	"strings"
)

// Equal compares two values and returns true if they are equal.
func Equal[T any](a, b T) bool {
	return reflect.DeepEqual(a, b)
}

// Kind returns true if the value is of the given kind.
func Kind(a any, kind reflect.Kind) bool {
	return reflect.TypeOf(a).Kind() == kind
}

// Nil returns true if the value is nil.
func Nil(a any) bool {
	if a == nil {
		return true
	}

	switch reflect.ValueOf(a).Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return reflect.ValueOf(a).IsNil()
	}

	return false
}

// Number returns true if the value is a number.
func Number(a any) bool {
	if a == nil {
		return false
	}

	switch reflect.TypeOf(a).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return true
	}
	return false
}

// Zero returns true if the value is the zero value.
func Zero(a any) bool {
	return a == nil || reflect.DeepEqual(a, reflect.Zero(reflect.TypeOf(a)).Interface())
}

// Implements returns true if the value implements the interface.
func Implements(a, iface any) bool {
	interfaceType := reflect.TypeOf(iface).Elem()

	if a == nil {
		return false
	}
	if !reflect.TypeOf(a).Implements(interfaceType) {
		return false
	}

	return true
}

// Panic returns true if the function panics.
func Panic(f func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	f()
	return
}

// Unique returns true if the slice contains unique values.
// Items are considered unique if they are not deep equal.
func Unique[T any](s []T) bool {
	// Use reflection DeepEqual to compare values.
	// This is necessary because slices of interfaces cannot be compared with ==.
	for i := range s {
		for j := range s {
			if i != j && reflect.DeepEqual(s[i], s[j]) {
				return false
			}
		}
	}
	return true
}

// Contains returns true if a contains b.
func Contains(a any, b any) bool {
	objectValue := reflect.ValueOf(a)
	objectKind := reflect.TypeOf(a).Kind()

	switch objectKind {
	case reflect.String:
		return strings.Contains(objectValue.String(), reflect.ValueOf(b).String())
	case reflect.Map:
	default:
		for i := 0; i < objectValue.Len(); i++ {
			if Equal(objectValue.Index(i).Interface(), b) {
				return true
			}
		}
	}

	return false
}

// ContainsAll returns true if all values are contained in a.
func ContainsAll[T any](a T, v []T) bool {
	for _, t := range v {
		if !Contains(a, t) {
			return false
		}
	}

	return true
}

// ContainsAny returns true if any of the values are contained in a.
func ContainsAny[T any](a T, v []T) bool {
	for _, t := range v {
		if Contains(a, t) {
			return true
		}
	}

	return false
}

// ContainsNone returns true if none of the values are contained in a.
func ContainsNone[T any](a T, v []T) bool {
	for _, t := range v {
		if Contains(a, t) {
			return false
		}
	}

	return true
}

// Uppercase returns true if the string is uppercase.
func Uppercase(s string) bool {
	return s == strings.ToUpper(s)
}

// Lowercase returns true if the string is lowercase.
func Lowercase(s string) bool {
	return s == strings.ToLower(s)
}
