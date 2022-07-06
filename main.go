package main

import (
	"fmt"
	"log"
	"payment-interface/payment"
)

func main() {
	vendor := payment.NewXendit()
	paymentService := payment.NewPayment(vendor)

	pay, err := paymentService.Pay()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pay)
}
