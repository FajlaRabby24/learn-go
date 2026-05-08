package main // executable package

import (
	"learngo_package/payment"
	// other "learngo_package/payment"  // * type alias
	"learngo_package/test"
)

func main() {
	// bkash := Bkash{apiKey: "932842lkjsdlfs"}
	// paymentService := PaymentService{
	// 	method: bkash,
	// }
	// paymentService := NewPaymentService(bkash)
	// paymentService.checkout()

	// nagad := Nagad{apiKey: "932842lkjsdlfs"}
	nagad := payment.NewNagad("932842lkjsdlfs")
	paymentService1 := payment.NewPaymentService(nagad)
	// paymentService1.checkout() // * small character was not exported
	paymentService1.Checkout()

	mockPm := test.MockPaymentMethod{}
	paymentService2 := payment.NewPaymentService(mockPm)
	paymentService2.Checkout()
}

// module -> bunch of packages
