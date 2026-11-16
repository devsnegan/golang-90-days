
package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}
type SMSNotifier struct{}

func (e EmailNotifier) Send(message string) {
	fmt.Println("Email sent:", message)
}

func (s SMSNotifier) Send(message string) {
	fmt.Println("SMS sent:", message)
}

func Notify(n Notifier, message string) {
	n.Send(message)
}

func main() {
	Notify(EmailNotifier{}, "Welcome!")
	Notify(SMSNotifier{}, "Your OTP is 1234")
}
