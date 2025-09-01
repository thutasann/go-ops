package howtousecomposeinterfaces

import (
	"fmt"
	"sync"
)

// Base logging behavior
type Logger interface {
	Log(message string)
}

// Formatter behavior
type Formatter interface {
	Format(message string) string
}

// Persister behavior
type Persister interface {
	Save(message string)
}

// Composed interface
type AdvancedLogger interface {
	Logger
	Formatter
	Persister
}

// Concrete implementation
type FileLogger struct {
	mu   sync.Mutex
	logs []string
}

func (f *FileLogger) Log(message string) {
	formatted := f.Format(message)
	f.Save(formatted)
}

func (f *FileLogger) Format(message string) string {
	return fmt.Sprintf("[Formatted] %s", message)
}

func (f *FileLogger) Save(message string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.logs = append(f.logs, message)
	fmt.Println("Saved:", message)
}

func Logger_Sample() {
	fmt.Println("\n===> Composed interface Logger Sample")

	var logger AdvancedLogger = &FileLogger{}
	logger.Log("System started")
	logger.Log("User logged in")
}
