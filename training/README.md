

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

- Pointers

```go
	i, j := 42, 207
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // devide j through the pointer
	fmt.Println(j) // see the new value of j
```

- Structs

```go
type Vertex struct {
	X, Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```

- Pointers to structs

```go
type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
```

- Struct Literals

```go
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{Y: 2}
	v3 = Vertex{}
	p  = &Vertex{3, 4}
)
```

- Arrays

```go
	var a [2]string
	a[0] = "Hello"
	a[1] = "world"
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)	
```

- Slices

```go
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	s := primes[1:3]
	fmt.Println(s)
```

- Slices are like references to arrays

```go
	a := [4]string{"AAA", "BBB", "CCC", "DDD"}
	b := a[0:2]
	c := a[1:3]
	fmt.Println(a, b, c)

	b[0] = "XXX"
	fmt.Println(a, b, c)
```

- Slice literals

```go
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)

	b := []bool{true, false, true, false}
	fmt.Println(b)

	c := []struct {
		x int
		y bool
	}{
		{1, true},
		{2, false},
		{3, true},
	}
	fmt.Println(c)
```

- Slice length and capacity

```go
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	s := []int{0, 2, 3, 4, 5}
	printSlice(s)

	// slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop it first two values.
	s = s[2:]
	printSlice(s)
}
```

- `nil` Slice

```go
	if s := []int; s == nil {
		fmt.Printf("Nil. len=%d cap=%d %v\n", len(s), cap(s), s)
	}
```

- Create a slice with `make`

```go
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
func main() {

	a := make([]int, 6)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[0:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}
```

- `slice` of slice

```go
	board := [][]string{
		{"_", "_", "_", "_"},
		{"_", "_", "_", "_"},
		{"_", "_", "_", "_"},
		{"_", "_", "_", "_"},
	}

	board[0][1] = "X"
	board[1][2] = "O"
	board[3][3] = "X"
	board[2][1] = "O"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
```

- `append`ing to a slice

```go
	var s []int
	s = append(s, 0)
	s = append(s, 1, 2, 3, 4)
```

- `range`

```go
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
```

- `map`

```go
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{1000.0, 2000.0}
	fmt.Println(m["Bell Labs"])

	m1 := map[string]Vertex{
		"Vietnix": {1.1, 2.2},
		"Google":  {3.3, 4.4},
	}
	fmt.Println(m1)

	delete(m1, "Google")
	fmt.Println(m1)

	v, ok := m1["Vietnix"]
	fmt.Println("The Value:", v, ". Present?", ok)
```

- Function values

