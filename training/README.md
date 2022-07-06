

# Section 01 - Golang syntax

- Functions

```go
func add(x, y int) {
    return x + y
}
```

- Multiple results

```go
func swap(x, y string ) (string, string) {
    return y, x
}
```

- Named return values

```go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```

- Variables

```go
var c, python, java pool
```

- Variables with initializers

```go
	var i, j int = 1, 2
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
```

- Short variable declarations

```go
	i, j := 1, 2
	c, python, java := true, false, "no!"
```

- Basic types

```go
bool

string

int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64

byte // alias for int8

rune // alias for int32
    // represents a Unicode point

float32 float64

complex64 complex128
```




