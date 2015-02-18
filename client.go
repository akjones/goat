package main

var msg = "something"

// GetMessage gets a message.
func GetMessage() string {
	return msg
}

// SendMessage sends a message.
func SendMessage(m string) error {
	msg = m
	return nil
}
