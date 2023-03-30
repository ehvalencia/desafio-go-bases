package main

import (
	"fmt"

	"github.com/ehvalencia/EjemplosGo/internal/tickets"
)

func main() {
	totalTickets, err := tickets.GetTotalTickets("Indonesia")
	if err != nil {
		panic(err)
	}
	fmt.Printf("El total de boletos vendidos para China es: %d\n", totalTickets)

}
