package howtousecomposeinterfaces

import "fmt"

// Base interface
type Service interface {
	FetchData(id string) (string, error)
}

// Implementation
type APIService struct{}

func (a *APIService) FetchData(id string) (string, error) {
	return "Data for " + id, nil
}

// Logging decorator (also implements Service)
type LoggingService struct {
	next Service
}

func (l *LoggingService) FetchData(id string) (string, error) {
	fmt.Println("Logging: fetching", id)
	return l.next.FetchData(id)
}

type CachingService struct {
	next  Service
	cache map[string]string
}

func (c *CachingService) FetchData(id string) (string, error) {
	if val, ok := c.cache[id]; ok {
		fmt.Println("cache hit: ", id)
		return val, nil
	}
	data, err := c.next.FetchData(id)
	if err == nil {
		c.cache[id] = data
	}
	return data, err
}

func Middleware_Pipeline_Pattern() {
	fmt.Println("\n===> Middleware_Pipeline Pattern")

	base := &APIService{}
	cached := &CachingService{next: base, cache: make(map[string]string)}
	logged := &LoggingService{next: cached}

	var service Service = logged

	fmt.Println(service.FetchData("123")) // logs + fetch
	fmt.Println(service.FetchData("123")) // logs + cache
}
