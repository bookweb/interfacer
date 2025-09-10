# interfacer

Generate accessor methods and interface for struct in Golang

## Installation

## Usage

### Declare Struct with `interfacer` Tag

`interfacer` generates accessor methods from defined structs, so you need to declare a struct and fields with `interfacer` tag.

Values for `interfacer` tag is `getter` and `setter`, `getter` is for generating getter method and `setter` is for setter methods.

Here is an example:

```go
type MyStruct struct {
    field1 string    `interfacer:"getter"`
    field2 *int      `interfacer:"setter"`
    field3 time.Time `interfacer:"getter,setter"`
}
```

Generated methods will be
```go
func(m *MyStruct) Field1() string {
    if m == nil {
        return ""
    }
    return m.field1
}

func(m *MyStruct) SetField2(val *int) {
    if m == nil {
        return
    }
    m.field2 = val
}

func(m *MyStruct) Field3() time.Time {
    if m == nil {
        return time.Time{}
    }
    return m.field3
}

func(m *MyStruct) SetField3(val time.Time) {
    if m == nil {
        return
    }
    m.field3 = val
}

type IMyStruct interface {
	Field1() string

	SetField2(val *int)

	Field3() string

	SetField3(val *string)
}
```

Following to [convention](https://golang.org/doc/effective_go#Getters),
setter's name is `Set<FieldName>()` and getter's name is `<FieldName>()` by default,
in other words, `Set` will be put into setter's name and `Get` will **not** be put into getter's name.

You can customize names for setter and getter if you want.

```go
type MyStruct struct {
    field1 string `interfacer:"getter:GetFirstField"`
    field2 int    `interfacer:"setter:ChangeSecondField"`
}
```

Generated methods will be

```go
func(m *MyStruct) GetFirstField() string {
    if m == nil {
        return ""
    }
    return m.field1
}

func(m *MyStruct) ChangeSecondField(val *int) {
    if m == nil {
        return
    }
    m.field2 = val
}
```

Accessor methods won't be generated if `interfacer` tag isn't specified.
But you can explicitly skip generation by using `-` for tag value.

```go
type MyStruct struct {
    ignoredField `interfacer:"-"`
}
```

### Specify rules for setting value
You can specify validation rules for each fields.
We underlying the 

### Run `interfacer` command

To generate accessor methods, you need to run `interfacer` command.

```
$ interfacer [flags] source-dir

source-dir
  source-dir is the directory where the definition of the target struct is located.
  If source-dir is not specified, current directory is set as source-dir.

flags
  --type string <required>
      name of target struct

  --receiver string <optional>
      receiver receiver for generated accessor methods
      default: first letter of struct

  --output string <optional>
      output file name
      default: <type_name>.gen.go

  --lock string <optional>
      specify lock field name and generate codes obtaining and releasing lock
      this is used to prevent race condition when concurrent access can be expected

  --version
      show the current version of interfacer
```

Example:

```shell
$ interfacer generate --type MyStruct --receiver myStruct --output my_struct.gen.go path/to/target
```

#### go generate

You can also generate accessors by using `go generate`.

```go
package mypackage

//go:generate interfacer -type MyStruct -receiver myStruct -output my_struct.gen.go

type MyStruct struct {
    field1 string `interfacer:"getter"`
    field2 *int   `interfacer:"setter"`
}
```

Then run go generate for your package.

## Credit 

<!-- Code generate accessor functions based on [accessory](https://github.com/masaushi/accessory){target=_blank} -->
Code generate accessor functions based on <a href="https://github.com/masaushi/accessory" target="_blank">accessory</a>

## License
