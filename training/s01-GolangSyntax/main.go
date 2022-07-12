package main

import (
	"errors"
	"fmt"
	"os"
)

var X, Y float64

var operator = map[string]func() (error, float64){
	// phep cong
	"1": func() (error, float64) {
		return nil, X + Y
	},
	// phep tru
	"2": func() (error, float64) {
		return nil, X - Y
	},
	// phep Nhan
	"3": func() (error, float64) {
		return nil, X * Y
	},
	// phep chia
	"4": func() (error, float64) {
		if Y == 0 {
			return errors.New("Khong the chia cho 0."), 0
		}
		return nil, X / Y
	},
}

func compute(fn func() (error, float64)) (error, float64) {
	return fn()
}

func inputValidate(v interface{}) (bool, interface{}) {

	switch true {
	case func() bool {
		_, ok := v.(float64)
		return ok
	}():
		return true, v.(float64)
	case func() bool {
		_, ok := v.(int)
		return ok
	}():
		return true, v.(int)

	}
	return false
}

func main() {

	var x, y interface{} = 23.0, 3

	if !inputValidate(x) {
		os.Exit(1)
	}
	if inputValidate(y) {
		os.Exit(2)
	}

	X = float64(x.(float64))
	Y = float64(y.(float64))

	if f, ok := operator["5"]; ok {
		fmt.Println(compute(f))
	} else {
		fmt.Println("Khong ho tro!")
	}

	fmt.Printf("%T", X)
}
