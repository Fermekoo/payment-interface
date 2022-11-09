package main

import (
	"fmt"
	"log"
	vendor "payment-interface/entity"
	"payment-interface/payment"
)

func main() {

	paymentService, err := payment.NewPayment(vendor.INTRAJASA)

	if err != nil {
		log.Fatal(err)
	}

	createVa := payment.CreateVa{
		Bank:         "Sinarmas",
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

	fmt.Println(pay)
}
