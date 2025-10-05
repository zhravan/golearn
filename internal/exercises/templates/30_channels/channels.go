package channels

import "fmt"

// TODO: 1. Create a channel
// TODO: 2. Send the senderMessage to main goroutine using the channel
// instead of printing it
// TODO: 3. Return the senderMessage from the function
func ReadMessage() string {
	go func() {
		senderMessage := "Hi!"
		fmt.Println(senderMessage)
	}()

	return ""
}
