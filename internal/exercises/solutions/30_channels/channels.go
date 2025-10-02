package channels

func ReadMessage() string {
	chat := make(chan string)
	go func() {
		senderMessage := "Hi!"
		chat <- senderMessage
	}()

	msg := <-chat

	return msg
}
