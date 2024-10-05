package main

import (
	"fmt"
	"io"
	"os"
)

// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

// func FPrintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
// 	p := newPrinter()
// 	p.doPrintf(format, a)
// 	n, err = w.Write(p.buf)
// 	p.free()
// 	return
// }

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Chetan")
}
