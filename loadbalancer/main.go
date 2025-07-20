package main

// Simple LoadBalancer in GoLang
func main() {
	servers := []Server{
		newSimpleServer("https://www.facebook.com/"),
		newSimpleServer("https://www.bing.com/"),
		newSimpleServer("https://www.duckduckgo.com/"),
	}

	lb := NewLoadBalancer("8000", servers)
}
