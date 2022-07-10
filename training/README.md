

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

- zero values

- Type conversions

```go
        var x, y int = 3, 4
        var f float64 = math.Sqrt(float64(x*x + y*y))
        var z uint = uint(f)
        fmt.Println(x, y, z, f)
```

- Type inference

```go
i := 1
f := 3.14
```

- Contants

```go
const Pi = 3.14

const (
	Big = 1 << 100
	Small = Big >> 99
)
```

- `For` loop

```go
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
```



- `For` continued

```go
	for ; i < 10 ; i++ {
```

- `For` is Go's "while"

```go
	for i < 10 {
```

- `For`ever

```go
	for {
	}
```

- `if`

```go
func sqrt(i float64) string {
	if i < 0 {
		return fmt.Sprint(math.Sqrt(-i)) + "i"
	}
	return fmt.Sprint(math.Sqrt(i))
}
```

- `if` with a short statement

```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		return lim
	}
}
```

- `switch`

```go
	fmt.Print("go run on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MAC OS")
	case "linux":
		fmt.Println("Linux")
	default:
		// BSD
		// Windows, other...
		fmt.Printf("%s\n", os)
	}
```

- `switch` evalution order

```go
	fmt.Println("When's Tuesday?")
	today := time.Now().Weekday()
	switch time.Tuesday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorow.")
	case today + 2:
		fmt.Println("In 2 days.")
	default:
		fmt.Println("Too far away.")
	}	
```

- `switch` with no condition

```go
	switch t := time.Now().Hour(); {
	case t < 12:
		fmt.Println("Good morning!")
	case t < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
```

- Stacking `defer`s

```go
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
```















