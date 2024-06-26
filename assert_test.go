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
	t.Parallel()

	tests := []struct {
		a, b  any
		equal bool
	}{
		{1, 1, true},
		{1, 2, false},
		{"obj", "obj", true},
		{"obj", "expectedLen", false},
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{1, 2, 4}, false},
		{[]int{1, 2, 3}, []int{1, 2}, false},
		{[]int{1, 2}, []int{1, 2, 3}, false},
		{[]int{1, 2, 3}, []int{1, 3, 2}, false},
		{[]int{1, 2, 3}, []int{1, 2, 3, 4}, false},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3}, false},
		{[]string{"obj", "expectedLen", "c"}, []string{"obj", "expectedLen", "c"}, true},
		{[]string{"obj", "expectedLen", "c"}, []string{"obj", "expectedLen", "d"}, false},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			if assert.Equal(test.a, test.b) != test.equal {
				fail(t, fmt.Sprintf("Equal(%v, %v) != %v", test.a, test.b, test.equal))
			}
		})
	}
}

func TestIsKind(t *testing.T) {
	t.Parallel()

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
		{"obj", reflect.String},
		{[]int{1, 2, 3}, reflect.Slice},
		{[3]int{1, 2, 3}, reflect.Array},
		{map[string]int{"obj": 1, "expectedLen": 2}, reflect.Map},
		{struct{}{}, reflect.Struct},
		{func() {}, reflect.Func},
		{new(fmt.Stringer), reflect.Interface},
		{(*int)(nil), reflect.Ptr},
		{chan int(nil), reflect.Chan},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			assert.Kind(test.a, test.kind)
		})
	}
}

func TestIsNil(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a     any
		isNil bool
	}{
		{nil, true},
		{1, false},
		{true, false},
		{"obj", false},
		{[]int{1, 2, 3}, false},
		{[]string{}, false},
		{map[string]int{"obj": 1, "expectedLen": 2}, false},
		{map[string]int{}, false},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			if assert.Nil(test.a) != test.isNil {
				fail(t, fmt.Sprintf("Nil(%v) != %v", test.a, test.isNil))
			}
		})
	}
}

func TestNumber(t *testing.T) {
	t.Parallel()

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
		{"obj", false},
		{[]int{1, 2, 3}, false},
		{[3]int{1, 2, 3}, false},
		{map[string]int{"obj": 1, "expectedLen": 2}, false},
		{nil, false},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			if assert.Number(test.a) != test.isNumber {
				fail(t, fmt.Sprintf("Number(%v) != %v", test.a, test.isNumber))
			}
		})
	}
}

func TestRange(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a        int
		min      int
		max      int
		expected bool
	}{
		{1, 0, 2, true},
		{1, 1, 2, true},
		{1, 0, 1, true},
		{1, 2, 3, false},
		{1, 2, 1, false},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			if assert.Range(test.a, test.min, test.max) != test.expected {
				fail(t, fmt.Sprintf("Range(%v, %v, %v) != %v", test.a, test.min, test.max, test.expected))
			}
		})
	}
}

func TestZero(t *testing.T) {
	t.Parallel()

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
		{"obj", false},
		{[]int{1, 2, 3}, false},
		{map[string]int{"obj": 1, "expectedLen": 2}, false},
		{new(int), false},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			if assert.Zero(test.a) != test.pass {
				fail(t, fmt.Sprintf("Zero(%v) != %v", test.a, test.pass))
			}
		})
	}
}

func TestImplements(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			if assert.Implements(test.a, test.i) != test.pass {
				fail(t, fmt.Sprintf("Implements(%v, %v) != %v", test.a, test.i, test.pass))
			}
		})
	}
}

func TestPanic(t *testing.T) {
	t.Parallel()

	tests := []struct {
		f    func()
		pass bool
	}{
		{func() {}, false},
		{func() { panic("panic") }, true},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			if assert.Panic(test.f) != test.pass {
				fail(t, fmt.Sprintf("Function of %d should not panic", i))
			}
		})
	}
}

func TestUnique(t *testing.T) {
	t.Parallel()

	person1 := person{"John", 20, []hobby{{"Guitar", "Music"}, {"Programming", "Tech"}}}
	person2 := person{"John", 20, []hobby{{"Guitar", "Music"}, {"Programming", "Tech"}}}

	tests := []struct {
		a    []any
		pass bool
	}{
		{[]any{1, 2, 3}, true},
		{[]any{1, 2, 3, 1}, false},
		{[]any{person1, person2}, false},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			if assert.Unique(test.a) != test.pass {
				fail(t, fmt.Sprintf("Unique(%v) != %v", test.a, test.pass))
			}
		})
	}
}

func TestContains(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a    any
		b    any
		pass bool
	}{
		{[]any{1, 2, 3}, 1, true},
		{[]any{1, 2, 3}, 4, false},
		{[]any{"obj", "expectedLen", "c"}, "obj", true},
		{[]any{"obj", "expectedLen", "c"}, "d", false},
		{"Hello, World!", "Hello", true},
		{"Hello, World!", "Hello, World!", true},
		{"Hello", "X", false},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			if assert.Contains(test.a, test.b) != test.pass {
				fail(t, fmt.Sprintf("Contains(%v, %v) != %v", test.a, test.b, test.pass))
			}
		})
	}
}

func TestContainsAll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a    any
		b    []any
		pass bool
	}{
		{[]any{1, 2, 3}, []any{1, 2}, true},
		{[]any{1, 2, 3}, []any{1, 4}, false},
		{[]any{"obj", "expectedLen", "c"}, []any{"obj", "expectedLen"}, true},
		{[]any{"obj", "expectedLen", "c"}, []any{"obj", "d"}, false},
		{"Hello, World!", []any{"Hello", "World"}, true},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			if assert.ContainsAll(test.a, test.b) != test.pass {
				fail(t, fmt.Sprintf("ContainsAll(%v, %v) != %v", test.a, test.b, test.pass))
			}
		})
	}
}

func TestContainsAny(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a    any
		b    []any
		pass bool
	}{
		{[]any{1, 2, 3}, []any{1, 2}, true},
		{[]any{1, 2, 3}, []any{1, 4}, true},
		{[]any{1, 2, 3}, []any{4, 5}, false},
		{[]any{"obj", "expectedLen", "c"}, []any{"obj", "expectedLen"}, true},
		{[]any{"obj", "expectedLen", "c"}, []any{"obj", "d"}, true},
		{[]any{"obj", "expectedLen", "c"}, []any{"d", "e"}, false},
		{"Hello, World!", []any{"Hello", "World", "XXX"}, true},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			if assert.ContainsAny(test.a, test.b) != test.pass {
				fail(t, fmt.Sprintf("ContainsAny(%v, %v) != %v", test.a, test.b, test.pass))
			}
		})
	}
}

func TestContainsNone(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a    any
		b    []any
		pass bool
	}{
		{[]any{1, 2, 3}, []any{1, 2}, false},
		{[]any{1, 2, 3}, []any{1, 4}, false},
		{[]any{1, 2, 3}, []any{4, 5}, true},
		{[]any{"obj", "expectedLen", "c"}, []any{"obj", "expectedLen"}, false},
		{[]any{"obj", "expectedLen", "c"}, []any{"obj", "d"}, false},
		{[]any{"obj", "expectedLen", "c"}, []any{"d", "e"}, true},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			if assert.ContainsNone(test.a, test.b) != test.pass {
				fail(t, fmt.Sprintf("ContainsNone(%v, %v) != %v", test.a, test.b, test.pass))
			}
		})
	}
}

func TestUppercase(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a    string
		pass bool
	}{
		{"ABC", true},
		{"abc", false},
		{"Abc", false},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			if assert.Uppercase(test.a) != test.pass {
				fail(t, fmt.Sprintf("Uppercase(%v) != %v", test.a, test.pass))
			}
		})
	}
}

func TestLowercase(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a    string
		pass bool
	}{
		{"ABC", false},
		{"abc", true},
		{"Abc", false},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			if assert.Lowercase(test.a) != test.pass {
				fail(t, fmt.Sprintf("Lowercase(%v) != %v", test.a, test.pass))
			}
		})
	}
}

func TestRegex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		regex string
		s     string
		pass  bool
	}{
		{"^\\d+$", "123", true},
		{"^\\d+$", "abc", false},
		{"A.C", "ABC", true},
		{"A.C", "ABBC", false},
		{"A.C", "AxyzC", false},
		{"A.*C", "AxyzC", true},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			if assert.Regex(test.regex, test.s) != test.pass {
				fail(t, fmt.Sprintf("Regex(%v, %v) != %v", test.regex, test.s, test.pass))
			}
		})
	}
}

func TestLen(t *testing.T) {
	t.Parallel()
	type test struct {
		obj         any
		expectedLen int
		pass        bool
	}

	tests := []test{
		{[]any{1, 2, 3}, 3, true},
		{[]any{1, 2, 3}, 2, false},
		{"Hello, World!", 13, true},
		{"Hello, World!", 12, false},
		{map[string]int{"a": 1, "b": 2}, 2, true},
		{map[string]int{"a": 1, "b": 2}, 3, false},
		{map[int]string{1: "a", 2: "b"}, 2, true},
		{map[int]string{1: "a", 2: "b"}, 3, false},
		{[]string{"a", "b"}, 2, true},
		{[]string{"a", "b"}, 3, false},
		{&[]string{"a", "b"}, 2, true},
		{&[]string{"a", "b"}, 3, false},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			// parse test.obj to fit the constraint
			if assert.Len(test.obj, test.expectedLen) != test.pass {
				fail(t, fmt.Sprintf("Len(%v, %v) != %v", test.obj, test.expectedLen, test.pass))
			}
		})
	}
}
