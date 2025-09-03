package howtousecomposeinterfaces

import "fmt"

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// Composed Reader Writer interface
type ReadWriter interface {
	Reader
	Writer
}

type File struct{}

func (f *File) Read(p []byte) (int, error) {
	fmt.Println("Reading data...")
	return len(p), nil
}

func (f *File) Write(p []byte) (int, error) {
	fmt.Println("Writing data...")
	return len(p), nil
}

func Reader_Writer_Sample() {
	fmt.Println("\n===> Reader Writer Sample")
	var rw ReadWriter = &File{}
	rw.Read([]byte("hello"))
	rw.Write([]byte("world"))
}
