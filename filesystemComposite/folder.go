package main

import "fmt"

type folder struct {
	name  string
	files []filesystem
}

func (f *folder) search(s string) {
	fmt.Printf("searching recursively for keyword %q in folder %q\n", s, f.name)
	for _, file := range f.files {
		file.search(s)
	}
}

func (f *folder) getName() string {
	return f.getName()
}

func (f *folder) add(file filesystem) {
	f.files = append(f.files, file)
}

/*
func (f *folder) add(files []filesystem) {
	for _, file := range files {
		f.files = append(f.files, file)
	}
}
*/
