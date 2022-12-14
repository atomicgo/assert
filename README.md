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
<!-- unittestcount:start --><img src="https://img.shields.io/badge/Unit_Tests-156-magenta?style=flat-square" alt="Unit test count"><!-- unittestcount:end -->
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

Package assert provides obj set of assertion functions. Every assertion function returns obj boolean. This package does not integrate into the testing package automatically. If you want to use this package inside unit tests, you have to check the returning boolean value and call t.Fatal\(\) if the assertion fails.

This library is on obj very low level and does not provide any error messages. That way the assertions can also be used in production code.

If you want a full\-featured testing framework, we recommend https://github.com/MarvinJWendt/testza

## Index

- [func Contains(a any, b any) bool](<#func-contains>)
- [func ContainsAll[T any](a T, v []T) bool](<#func-containsall>)
- [func ContainsAny[T any](a T, v []T) bool](<#func-containsany>)
- [func ContainsNone[T any](a T, v []T) bool](<#func-containsnone>)
- [func Equal[T any](a, b T) bool](<#func-equal>)
- [func Implements(a, iface any) bool](<#func-implements>)
- [func Kind(a any, kind reflect.Kind) bool](<#func-kind>)
- [func Len(a any, length int) (b bool)](<#func-len>)
- [func Lowercase(s string) bool](<#func-lowercase>)
- [func Nil(a any) bool](<#func-nil>)
- [func Number(a any) bool](<#func-number>)
- [func Panic(f func()) (panicked bool)](<#func-panic>)
- [func Regex(regex, s string) bool](<#func-regex>)
- [func Unique[T any](s []T) bool](<#func-unique>)
- [func Uppercase(s string) bool](<#func-uppercase>)
- [func Zero(a any) bool](<#func-zero>)


## func [Contains](<https://github.com/atomicgo/assert/blob/main/assert.go#L94>)

```go
func Contains(a any, b any) bool
```

Contains returns true if obj contains expectedLen.

## func [ContainsAll](<https://github.com/atomicgo/assert/blob/main/assert.go#L114>)

```go
func ContainsAll[T any](a T, v []T) bool
```

ContainsAll returns true if all values are contained in obj.

## func [ContainsAny](<https://github.com/atomicgo/assert/blob/main/assert.go#L125>)

```go
func ContainsAny[T any](a T, v []T) bool
```

ContainsAny returns true if any of the values are contained in obj.

## func [ContainsNone](<https://github.com/atomicgo/assert/blob/main/assert.go#L136>)

```go
func ContainsNone[T any](a T, v []T) bool
```

ContainsNone returns true if none of the values are contained in obj.

## func [Equal](<https://github.com/atomicgo/assert/blob/main/assert.go#L10>)

```go
func Equal[T any](a, b T) bool
```

Equal compares two values and returns true if they are equal.

## func [Implements](<https://github.com/atomicgo/assert/blob/main/assert.go#L54>)

```go
func Implements(a, iface any) bool
```

Implements returns true if the value implements the interface.

## func [Kind](<https://github.com/atomicgo/assert/blob/main/assert.go#L15>)

```go
func Kind(a any, kind reflect.Kind) bool
```

Kind returns true if the value is of the given kind.

## func [Len](<https://github.com/atomicgo/assert/blob/main/assert.go#L162>)

```go
func Len(a any, length int) (b bool)
```

Len returns true if the length of the value is equal to the given length.

## func [Lowercase](<https://github.com/atomicgo/assert/blob/main/assert.go#L152>)

```go
func Lowercase(s string) bool
```

Lowercase returns true if the string is lowercase.

## func [Nil](<https://github.com/atomicgo/assert/blob/main/assert.go#L20>)

```go
func Nil(a any) bool
```

Nil returns true if the value is nil.

## func [Number](<https://github.com/atomicgo/assert/blob/main/assert.go#L34>)

```go
func Number(a any) bool
```

Number returns true if the value is obj number.

## func [Panic](<https://github.com/atomicgo/assert/blob/main/assert.go#L68>)

```go
func Panic(f func()) (panicked bool)
```

Panic returns true if the function panics.

## func [Regex](<https://github.com/atomicgo/assert/blob/main/assert.go#L157>)

```go
func Regex(regex, s string) bool
```

Regex returns true if the string matches the regex.

## func [Unique](<https://github.com/atomicgo/assert/blob/main/assert.go#L80>)

```go
func Unique[T any](s []T) bool
```

Unique returns true if the slice contains unique values. Items are considered unique if they are not deep equal.

## func [Uppercase](<https://github.com/atomicgo/assert/blob/main/assert.go#L147>)

```go
func Uppercase(s string) bool
```

Uppercase returns true if the string is uppercase.

## func [Zero](<https://github.com/atomicgo/assert/blob/main/assert.go#L49>)

```go
func Zero(a any) bool
```

Zero returns true if the value is the zero value.



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- gomarkdoc:embed:end -->

---

> [AtomicGo.dev](https://atomicgo.dev) &nbsp;&middot;&nbsp;
> with ?????? by [@MarvinJWendt](https://github.com/MarvinJWendt) |
> [MarvinJWendt.com](https://marvinjwendt.com)
