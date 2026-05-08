package main // executable package

import (
	"learngo_package/payment"
	// other "learngo_package/payment"  // * type alias
	"learngo_package/test"

	"github.com/fatih/color"
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

	color.Red("Prints text in cyan.")
	color.BgRGB(255, 128, 0).Println("Background orange")
}

// module -> bunch of packages
