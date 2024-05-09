package main

import (
	"fmt"

	"github.com/hurtsec/fem-go-aws/imports"
)

func main() {
	newTicket := imports.Ticket{
		ID:    123,
		Event: "FEM",
	}

	newTicket.PrintEvent()

	fmt.Println(newTicket)
}
