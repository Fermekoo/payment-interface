package main

import (
	"fmt"
	"log"
	c "payment-interface/constants"
	"payment-interface/payment"
)

func main() {

	paymentService, err := payment.NewPayment(c.INTRAJASA) // xendit/mitrans/intrajasa/gopay

	if err != nil {
		log.Fatal(err)
	}

	createVa := payment.CreateVa{
		Bank:         "BRI",
		Name:         "Dandi Fermeko",
		Amount:       25000,
		Address:      "Jakarta Selatan",
		Email:        "dandifermeko@gmail.com",
		RegisterDate: "2020-01-01",
	}

	pay, err := paymentService.Pay(&createVa)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pay.OrderID)
}
