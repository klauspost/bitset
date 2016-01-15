# bitset
Go Integer Bitset Generator

Generator for bit sets stored in integers as custom types.

This will generate helper functions that will make it possible for you to handle bit sets easily in integers.

See the [Example godoc](https://godoc.org/github.com/klauspost/bitset/example) for typical output.

Package home: https://github.com/klauspost/bitset

# alpha warning

This software is very rough, and meant for a "generate once" scenario.

If there are things that needs adjusment in the generated code, you are encouraged to do so.

# usage

```
go get github.com/klauspost/bitset
cd %GOPATH%/src/github.com/klauspost/bitset
go run bitset.go MySet MyType example int >example/myset.go
```

* `MySet` is the name of the type you want to generate.
* `MyType` is the type name of your existing type.
* `example` is the name of the package.
* `int` is the underlying type of `MyType`. Must be convertible to uint64.
