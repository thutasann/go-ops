package icmethodsreceivers

import "fmt"

// Value receiver (func (t T)) → operates on a copy of the value.
// Pointer receiver (func (t *T)) → operates on the actual instance (can modify it).

// Value receiver = You’re working on a photo copy of a document.
// Pointer receiver = You’re working on the original document — edits persist.

// Counter
type Counter struct {
	Value int
}

// Method with a VALUE receiver
// This operates on a COPY of 'Counter'
func (c Counter) IncrementByCopy() {
	c.Value++
	fmt.Printf("[Value Receiver] Incremented copy: %d (original unchanged)\n", c.Value)
}

// Method with a POINTER receiver
// This operates on the ACTUAL instance
func (c *Counter) IncrementByPointer() {
	c.Value++
	fmt.Printf("[Pointer Receiver] Incremented original: %d (changes persist)\n", c.Value)
}

func Methods_Receivers() {
	fmt.Println("\n===> Methods Receivers")

	counter := Counter{Value: 10}

	fmt.Println("Initial Counter Value:", counter.Value)

	// 4️⃣ Call method with value receiver
	counter.IncrementByCopy()
	fmt.Println("After IncrementByCopy:", counter.Value)

	// 5️⃣ Call method with pointer receiver
	counter.IncrementByPointer()
	fmt.Println("After IncrementByPointer:", counter.Value)
}
