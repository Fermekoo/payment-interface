package main

import (
	"fmt"
	"log"
	"payment-interface/payment"
)

func main() {

	paymentService := payment.NewPayment("intrajasa") // xendit/mitrans/intrajasa

	createVa := payment.CreateVa{
		Bank:         "bca",
		Name:         "Dandi Fermeko",
		Amount:       25000,
		Address:      "Jakarta Selatan",
		Email:        "dandifermeko@gmail.com",
		RegisterDate: "2020-01-01",
	}

	pay, err := paymentService.Pay(createVa)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pay)
}
