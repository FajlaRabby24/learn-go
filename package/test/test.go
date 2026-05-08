package test

import "fmt"

type MockPaymentMethod struct {
}

func (mockPm MockPaymentMethod) Pay(amount float64) {
	fmt.Printf("Paying %.2f tk with Mock", amount)
}
