package tickets

import (
	"encoding/csv"
	"os"
)

type Ticket struct {
}

func GetTotalTickets(destination string) (int, error) {
	// abrir archivo csv
	file, err := os.Open("tickets.csv")
	if err != nil {
		return 0, err
	}
	//cerrar
	defer file.Close()

	// creao un lector csv que lea registros del archivo
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return 0, err
	}

	//creo una vatiabla para conar cuenato regiustros coiciden con el parametro que se pasa
	count := 0
	for _, record := range records {
		if record[3] == destination {
			count++
		}
	}
	//retorno el contador qie trae el total
	return count, nil
}

// ejemplo 2
//func GetMornings(time string) (int error) {}

// ejemplo 3
//func AverageDestination(destination string, total int) (int error) {}
