package main

import "fmt"

type file struct {
	name string
}

func (f *file) search(s string) {
	fmt.Printf("searching for keyword %q in file %q\n", s, f.name)
}

func (f *file) getName() string {
	return f.name
}
