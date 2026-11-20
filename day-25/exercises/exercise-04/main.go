package main

import "fmt"

type Notifier interface {
    Send()
}

type EmailNotifier struct{}

func (e EmailNotifier) Send() {
    fmt.Println("Sending email")
}

type SMSNotifier struct{}

func (s SMSNotifier) Send() {
    fmt.Println("Sending SMS")
}

func Notify(value any) {
    switch v := value.(type) {
    case EmailNotifier:
        fmt.Println("Email notifier")
        v.Send()

    case SMSNotifier:
        fmt.Println("SMS notifier")
        v.Send()

    default:
        fmt.Println("Unknown notifier")
    }
}

func main() {
    Notify(EmailNotifier{})
    Notify(SMSNotifier{})
}