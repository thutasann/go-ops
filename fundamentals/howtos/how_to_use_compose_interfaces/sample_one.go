package howtousecomposeinterfaces

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
)

// Hash Reader interface
type HashReader interface {
	io.Reader     // Reader
	hash() string // Hash fn
}

// hash reader struct
type hashReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

// Initialize new Hash Reader
func NewHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

// Impl hash fn
func (h *hashReader) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

// Hash and Broadcast
func hashAndBroadcast(r HashReader) error {
	hash := r.hash()
	fmt.Println("hash:", hash)
	return broadcast(r)
}

// Broadcast fn
func broadcast(r io.Reader) error {

	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	fmt.Println("string of the bytes:", string(b))

	return nil
}

func How_To_Use_Compose_Interfaces_Sample_One() {
	fmt.Println("\n==> How to use compose interfaces sample one")

	payload := []byte("hello high high value software engineer")
	hashAndBroadcast(NewHashReader(payload))
}
