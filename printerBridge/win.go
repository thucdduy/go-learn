package main

import "fmt"

type win struct {
	printer printer
}

func (w *win) setPrinter(p printer) {
	w.printer = p
}

func (w *win) print() {
	fmt.Println("print request for win")
	w.printer.printFile()

}
