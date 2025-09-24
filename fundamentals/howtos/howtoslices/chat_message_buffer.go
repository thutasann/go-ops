package howtoslices

import "fmt"

type Chat struct {
	history []string
	limit   int
}

func NewChat(limit int) *Chat {
	return &Chat{limit: limit}
}

func (c *Chat) Add(msg string) {
	if len(c.history) == c.limit {
		c.history = c.history[1:] // drop oldest
	}
	c.history = append(c.history, msg)
}

func (c *Chat) History() []string {
	return c.history
}

func Chat_Message_Buffer() {
	fmt.Println("\n==> Chat Message Buffer")

	chat := NewChat(3)

	chat.Add("Hello")
	chat.Add("How are you?")
	chat.Add("Fine thanks")
	chat.Add("What about you?")

	fmt.Println(chat.History()) // [How are you? Fine thanks What about you?]
}
