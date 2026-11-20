package main

import "fmt"

type CashPayment struct{}
type CardPayment struct{}
type UPIPayment struct{}

func ProcessPayment(payment any) {
    switch payment.(type) {
    case CashPayment:
        fmt.Println("Processing cash payment")
    case CardPayment:
        fmt.Println("Processing card payment")
    case UPIPayment:
        fmt.Println("Processing UPI payment")
    default:
        fmt.Println("Unsupported payment")
    }
}

func main() {
    ProcessPayment(CashPayment{})
    ProcessPayment(CardPayment{})
    ProcessPayment(UPIPayment{})
}