package main

import (
	"fmt"
	"log"
	"payment-interface/payment"
)

func main() {

	vendor := payment.NewMidtrans()
	paymentService := payment.NewPayment(vendor)

	createVa := payment.CreateVa{
		Bank:   "bca",
		Name:   "Dandi Fermeko",
		Amount: 25000,
	}

	pay, err := paymentService.Pay(createVa)

	if err != nil {
		log.Fatal(err)
	}

	inquiry, err := paymentService.Inquiry(pay.OrderID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pay)
	fmt.Println(inquiry)
}
