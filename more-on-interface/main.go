package main

import "fmt"

type PaymentMethod interface {
	pay(amount float64)
}

type Bkash struct {
	apiKey string
}
type Nagad struct {
	apiKey string
}

func (bk Bkash) pay(amount float64) {
	// actual payment login
	fmt.Printf("Paying %.2f tk with Bkash", amount)
}
func (bk *Nagad) pay(amount float64) {
	// actual payment login
	fmt.Printf("Paying %.2f tk with Nagad", amount)
}

type PaymentService struct {
	method PaymentMethod
}

func NewPaymentService(method PaymentMethod) *PaymentService {
	return &PaymentService{
		method: method,
	}
}

func NewNagad(apiKey string) *Nagad {
	return &Nagad{
		apiKey,
	}
}

func (ps PaymentService) checkout() {
	// bkash := Bkash{apiKey: "932842lkjsdlfs"}
	// bkash.pay(100.00)

	ps.method.pay(100.00)
}

type MockPaymentMethod struct {
}

func (mockPm MockPaymentMethod) pay(amount float64) {
	fmt.Printf("Paying %.2f tk with Mock", amount)
}

func main() {
	// bkash := Bkash{apiKey: "932842lkjsdlfs"}
	// paymentService := PaymentService{
	// 	method: bkash,
	// }
	// paymentService := NewPaymentService(bkash)
	// paymentService.checkout()

	// nagad := Nagad{apiKey: "932842lkjsdlfs"}
	nagad := NewNagad("932842lkjsdlfs")
	paymentService1 := NewPaymentService(nagad)
	paymentService1.checkout()

	mockPm := MockPaymentMethod{}
	paymentService2 := NewPaymentService(mockPm)
	paymentService2.checkout()
}
