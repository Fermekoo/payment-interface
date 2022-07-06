package main

import (
	"fmt"
	"log"
	"payment-interface/payment"
)

func main() {

	vendor := payment.NewXendit()
	paymentService := payment.NewPayment(vendor)

	createVa := payment.CreateVa{
		Bank:   "permata",
		Name:   "Dandi Fermeko",
		Amount: 25000,
	}
	pay, err := paymentService.Pay(createVa)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pay)
}
