<h1 align="center">AtomicGo | assert</h1>

<p align="center">
<img src="https://img.shields.io/endpoint?url=https%3A%2F%2Fatomicgo.dev%2Fapi%2Fshields%2Fassert&style=flat-square" alt="Downloads">

<a href="https://github.com/atomicgo/assert/releases">
<img src="https://img.shields.io/github/v/release/atomicgo/assert?style=flat-square" alt="Latest Release">
</a>

<a href="https://codecov.io/gh/atomicgo/assert" target="_blank">
<img src="https://img.shields.io/github/actions/workflow/status/atomicgo/assert/go.yml?style=flat-square" alt="Tests">
</a>

<a href="https://codecov.io/gh/atomicgo/assert" target="_blank">
<img src="https://img.shields.io/codecov/c/gh/atomicgo/assert?color=magenta&logo=codecov&style=flat-square" alt="Coverage">
</a>

<a href="https://codecov.io/gh/atomicgo/assert">
<!-- unittestcount:start --><img src="https://img.shields.io/badge/Unit_Tests-162-magenta?style=flat-square" alt="Unit test count"><!-- unittestcount:end -->
</a>

<a href="https://opensource.org/licenses/MIT" target="_blank">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>
  
<a href="https://goreportcard.com/report/github.com/atomicgo/assert" target="_blank">
<img src="https://goreportcard.com/badge/github.com/atomicgo/assert?style=flat-square" alt="Go report">
</a>   

</p>

---

<p align="center">
<strong><a href="https://pkg.go.dev/atomicgo.dev/assert#section-documentation" target="_blank">Documentation</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CONTRIBUTING.md" target="_blank">Contributing</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CODE_OF_CONDUCT.md" target="_blank">Code of Conduct</a></strong>
</p>

---

<p align="center">
  <img src="https://raw.githubusercontent.com/atomicgo/atomicgo/main/assets/header.png" alt="AtomicGo">
</p>

<p align="center">
<table>
<tbody>
</tbody>
</table>
</p>
<h3  align="center"><pre>go get atomicgo.dev/assert</pre></h3>
<p align="center">
<table>
<tbody>
</tbody>
</table>
</p>

<!-- gomarkdoc:embed:start -->

<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# assert

```go
import "atomicgo.dev/assert"
```

Package assert provides a set of assertion functions. Every assertion function returns a boolean. This package does not integrate into the testing package automatically. If you want to use this package inside unit tests, you have to check the returning boolean value and call t.Fatal\(\) if the assertion fails.

This library does not provide any error messages. That way the assertions can also be used in production code.

If you want a full\-featured testing framework, we recommend https://github.com/MarvinJWendt/testza \(which uses this library for assertions\)

## Index

- [func Contains\(a any, b any\) bool](<#Contains>)
- [func ContainsAll\[T any\]\(a T, v \[\]T\) bool](<#ContainsAll>)
- [func ContainsAny\[T any\]\(a T, v \[\]T\) bool](<#ContainsAny>)
- [func ContainsNone\[T any\]\(a T, v \[\]T\) bool](<#ContainsNone>)
- [func Equal\[T any\]\(a, b T\) bool](<#Equal>)
- [func Implements\(a, iface any\) bool](<#Implements>)
- [func Kind\(a any, kind reflect.Kind\) bool](<#Kind>)
- [func Len\(a any, length int\) \(b bool\)](<#Len>)
- [func Lowercase\(s string\) bool](<#Lowercase>)
- [func Nil\(a any\) bool](<#Nil>)
- [func Number\(a any\) bool](<#Number>)
- [func Panic\(f func\(\)\) \(panicked bool\)](<#Panic>)
- [func Range\[T constraints.Number\]\(a, min, max T\) bool](<#Range>)
- [func Regex\(regex, s string\) bool](<#Regex>)
- [func Unique\[T any\]\(s \[\]T\) bool](<#Unique>)
- [func Uppercase\(s string\) bool](<#Uppercase>)
- [func Zero\(a any\) bool](<#Zero>)


<a name="Contains"></a>
## func [Contains](<https://github.com/atomicgo/assert/blob/main/assert.go#L107>)

```go
func Contains(a any, b any) bool
```

Contains returns true if obj contains expectedLen.

<a name="ContainsAll"></a>
## func [ContainsAll](<https://github.com/atomicgo/assert/blob/main/assert.go#L127>)

```go
func ContainsAll[T any](a T, v []T) bool
```

ContainsAll returns true if all values are contained in obj.

<a name="ContainsAny"></a>
## func [ContainsAny](<https://github.com/atomicgo/assert/blob/main/assert.go#L138>)

```go
func ContainsAny[T any](a T, v []T) bool
```

ContainsAny returns true if any of the values are contained in obj.

<a name="ContainsNone"></a>
## func [ContainsNone](<https://github.com/atomicgo/assert/blob/main/assert.go#L149>)

```go
func ContainsNone[T any](a T, v []T) bool
```

ContainsNone returns true if none of the values are contained in obj.

<a name="Equal"></a>
## func [Equal](<https://github.com/atomicgo/assert/blob/main/assert.go#L12>)

```go
func Equal[T any](a, b T) bool
```

Equal compares two values and returns true if they are equal.

<a name="Implements"></a>
## func [Implements](<https://github.com/atomicgo/assert/blob/main/assert.go#L64>)

```go
func Implements(a, iface any) bool
```

Implements returns true if the value implements the interface.

<a name="Kind"></a>
## func [Kind](<https://github.com/atomicgo/assert/blob/main/assert.go#L17>)

```go
func Kind(a any, kind reflect.Kind) bool
```

Kind returns true if the value is of the given kind.

<a name="Len"></a>
## func [Len](<https://github.com/atomicgo/assert/blob/main/assert.go#L175>)

```go
func Len(a any, length int) (b bool)
```

Len returns true if the length of the value is equal to the given length.

<a name="Lowercase"></a>
## func [Lowercase](<https://github.com/atomicgo/assert/blob/main/assert.go#L165>)

```go
func Lowercase(s string) bool
```

Lowercase returns true if the string is lowercase.

<a name="Nil"></a>
## func [Nil](<https://github.com/atomicgo/assert/blob/main/assert.go#L22>)

```go
func Nil(a any) bool
```

Nil returns true if the value is nil.

<a name="Number"></a>
## func [Number](<https://github.com/atomicgo/assert/blob/main/assert.go#L37>)

```go
func Number(a any) bool
```

Number returns true if the value is obj number.

<a name="Panic"></a>
## func [Panic](<https://github.com/atomicgo/assert/blob/main/assert.go#L79>)

```go
func Panic(f func()) (panicked bool)
```

Panic returns true if the function panics.

<a name="Range"></a>
## func [Range](<https://github.com/atomicgo/assert/blob/main/assert.go#L54>)

```go
func Range[T constraints.Number](a, min, max T) bool
```

Range returns true if the value is within the range.

<a name="Regex"></a>
## func [Regex](<https://github.com/atomicgo/assert/blob/main/assert.go#L170>)

```go
func Regex(regex, s string) bool
```

Regex returns true if the string matches the regex.

<a name="Unique"></a>
## func [Unique](<https://github.com/atomicgo/assert/blob/main/assert.go#L92>)

```go
func Unique[T any](s []T) bool
```

Unique returns true if the slice contains unique values. Items are considered unique if they are not deep equal.

<a name="Uppercase"></a>
## func [Uppercase](<https://github.com/atomicgo/assert/blob/main/assert.go#L160>)

```go
func Uppercase(s string) bool
```

Uppercase returns true if the string is uppercase.

<a name="Zero"></a>
## func [Zero](<https://github.com/atomicgo/assert/blob/main/assert.go#L59>)

```go
func Zero(a any) bool
```

Zero returns true if the value is the zero value.

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- gomarkdoc:embed:end -->

---

> [AtomicGo.dev](https://atomicgo.dev) &nbsp;&middot;&nbsp;
> with ❤️ by [@MarvinJWendt](https://github.com/MarvinJWendt) |
> [MarvinJWendt.com](https://marvinjwendt.com)
