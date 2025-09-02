package howtousecomposeinterfaces

import "fmt"

type Middleware interface {
	Process(data string) string
}

type LoggerMiddleware struct{}

func (l LoggerMiddleware) Process(data string) string {
	fmt.Println("Logging: ", data)
	return data
}

type EncryptMiddleware struct{}

func (e EncryptMiddleware) Process(data string) string {
	return "[encrypted]" + data
}

// Composed
type Pipeline interface {
	Middleware
}

func run_pipeline(p []Pipeline, data string) string {
	for _, stage := range p {
		data = stage.Process(data)
	}
	return data
}

func Middleware_Pipeline() {
	fmt.Println("\n===> Composed interface Middleware Pipeline Sample")

	pipeline := []Pipeline{
		LoggerMiddleware{},
		EncryptMiddleware{},
	}

	result := run_pipeline(pipeline, "Hello World")
	fmt.Println("Final:", result)
}
