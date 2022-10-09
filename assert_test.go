package assert_test

import (
	"fmt"
	"go/types"
	"reflect"
	"strconv"
	"testing"

	"atomicgo.dev/assert"
)

func fail(t *testing.T, msg string) {
	t.Helper()
	t.Fatal(msg)
}

func TestEqual(t *testing.T) {
	tests := []struct {
		a, b  any
		equal bool
	}{
		{1, 1, true},
		{1, 2, false},
		{"a", "a", true},
		{"a", "b", false},
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{1, 2, 4}, false},
		{[]int{1, 2, 3}, []int{1, 2}, false},
		{[]int{1, 2}, []int{1, 2, 3}, false},
		{[]int{1, 2, 3}, []int{1, 3, 2}, false},
		{[]int{1, 2, 3}, []int{1, 2, 3, 4}, false},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3}, false},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}, true},
		{[]string{"a", "b", "c"}, []string{"a", "b", "d"}, false},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if assert.Equal(test.a, test.b) != test.equal {
				fail(t, fmt.Sprintf("Equal(%v, %v) != %v", test.a, test.b, test.equal))
			}
		})
	}
}

func TestIsKind(t *testing.T) {
	tests := []struct {
		a    any
		kind reflect.Kind
	}{
		{1, reflect.Int},
		{int8(1), reflect.Int8},
		{int16(1), reflect.Int16},
		{int32(1), reflect.Int32},
		{int64(1), reflect.Int64},
		{uint(1), reflect.Uint},
		{uint8(1), reflect.Uint8},
		{uint16(1), reflect.Uint16},
		{uint32(1), reflect.Uint32},
		{uint64(1), reflect.Uint64},
		{uintptr(1), reflect.Uintptr},
		{float32(1), reflect.Float32},
		{float64(1), reflect.Float64},
		{complex64(1), reflect.Complex64},
		{complex128(1), reflect.Complex128},
		{true, reflect.Bool},
		{"a", reflect.String},
		{[]int{1, 2, 3}, reflect.Slice},
		{[3]int{1, 2, 3}, reflect.Array},
		{map[string]int{"a": 1, "b": 2}, reflect.Map},
		{struct{}{}, reflect.Struct},
		{func() {}, reflect.Func},
		{new(fmt.Stringer), reflect.Interface},
		{(*int)(nil), reflect.Ptr},
		{chan int(nil), reflect.Chan},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Kind(test.a, test.kind)
		})
	}
}

func TestIsNil(t *testing.T) {
	tests := []struct {
		a     any
		isNil bool
	}{
		{nil, true},
		{1, false},
		{true, false},
		{"a", false},
		{[]int{1, 2, 3}, false},
		{[]string{}, false},
		{map[string]int{"a": 1, "b": 2}, false},
		{map[string]int{}, false},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if assert.Nil(test.a) != test.isNil {
				fail(t, fmt.Sprintf("Nil(%v) != %v", test.a, test.isNil))
			}
		})
	}
}

func TestNumber(t *testing.T) {
	tests := []struct {
		a        any
		isNumber bool
	}{
		{1, true},
		{int8(1), true},
		{int16(1), true},
		{int32(1), true},
		{int64(1), true},
		{uint(1), true},
		{uint8(1), true},
		{uint16(1), true},
		{uint32(1), true},
		{uint64(1), true},
		{uintptr(1), true},
		{float32(1), true},
		{float64(1), true},
		{complex64(1), true},
		{complex128(1), true},
		{true, false},
		{"a", false},
		{[]int{1, 2, 3}, false},
		{[3]int{1, 2, 3}, false},
		{map[string]int{"a": 1, "b": 2}, false},
		{nil, false},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if assert.Number(test.a) != test.isNumber {
				fail(t, fmt.Sprintf("Number(%v) != %v", test.a, test.isNumber))
			}
		})
	}
}

func TestTrue(t *testing.T) {
	tests := []struct {
		a    bool
		pass bool
	}{
		{true, true},
		{false, false},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if assert.True(test.a) != test.pass {
				fail(t, fmt.Sprintf("True(%v) != %v", test.a, test.pass))
			}
		})
	}
}

func TestFalse(t *testing.T) {
	tests := []struct {
		a    bool
		pass bool
	}{
		{true, false},
		{false, true},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if assert.False(test.a) != test.pass {
				fail(t, fmt.Sprintf("False(%v) != %v", test.a, test.pass))
			}
		})
	}
}

func TestZero(t *testing.T) {
	tests := []struct {
		a    any
		pass bool
	}{
		{0, true},
		{0.0, true},
		{0i, true},
		{false, true},
		{"", true},
		{nil, true},
		{1, false},
		{1.0, false},
		{1i, false},
		{true, false},
		{"a", false},
		{[]int{1, 2, 3}, false},
		{map[string]int{"a": 1, "b": 2}, false},
		{new(int), false},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if assert.Zero(test.a) != test.pass {
				fail(t, fmt.Sprintf("Zero(%v) != %v", test.a, test.pass))
			}
		})
	}
}

func TestImplements(t *testing.T) {
	tests := []struct {
		a    any
		i    interface{}
		pass bool
	}{
		{new(types.Const), (*fmt.Stringer)(nil), true},
		{new(types.Const), (*types.Type)(nil), false},
		{new(types.Const), (*types.Type)(nil), false},
		{new(types.Basic), (*types.Type)(nil), true},
		{new(types.Basic), (*fmt.Stringer)(nil), true},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if assert.Implements(test.a, test.i) != test.pass {
				fail(t, fmt.Sprintf("Implements(%v, %v) != %v", test.a, test.i, test.pass))
			}
		})
	}
}

func TestPanic(t *testing.T) {
	tests := []struct {
		f    func()
		pass bool
	}{
		{func() {}, false},
		{func() { panic("panic") }, true},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if assert.Panic(test.f) != test.pass {
				fail(t, fmt.Sprintf("Function of %d should not panic", i))
			}
		})
	}
}
