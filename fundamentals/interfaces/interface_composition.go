package interfaces

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct{}

func (f File) Read() string {
	return "data from file"
}

func (f File) Write(data string) {
	fmt.Println("writing...", data)
}

func Interface_Composition(rw ReadWriter) {
	data := rw.Read()
	rw.Write(data + "_processed")
}
