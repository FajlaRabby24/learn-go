package payment

import "fmt"

type PaymentMethod interface {
	Pay(amount float64)
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
func (bk *Nagad) Pay(amount float64) {
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

func (ps PaymentService) Checkout() {
	// bkash := Bkash{apiKey: "932842lkjsdlfs"}
	// bkash.pay(100.00)

	ps.method.Pay(100.00)
}
