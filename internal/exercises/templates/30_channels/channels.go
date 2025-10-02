package channels

import "fmt"

func ReadMessage() string {
	// TODO: update following code
	// to receive senderMessage in main goroutine
	// and return it from ReadMessage function
	go func() {
		senderMessage := "Hi!"
		fmt.Println(senderMessage)
	}()

	return ""
}
