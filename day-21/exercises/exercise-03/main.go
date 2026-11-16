
package main

import "fmt"

type PaymentMethod interface {
	Pay(amount float64)
}

type UPI struct{}
type CreditCard struct{}

func (u UPI) Pay(amount float64) {
	fmt.Println("UPI payment:", amount)
}

func (c CreditCard) Pay(amount float64) {
	fmt.Println("Credit card payment:", amount)
}

func ProcessPayment(p PaymentMethod, amount float64) {
	p.Pay(amount)
}

func main() {
	ProcessPayment(UPI{}, 1000)
	ProcessPayment(CreditCard{}, 2000)
}
