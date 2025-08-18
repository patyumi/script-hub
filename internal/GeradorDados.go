package internal

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
	"time"
)

func GeradorDados() {

	fmt.Println("Gerador de dados selecionado!")

	startTime := time.Now()

	nameSet := []string{"Alice", "Bob", "Charlie", "Diana"}
	stateSet := []string{"California", "Texas", "New York", "Florida"}
	header := []string{"Name", "Age", "State"}
	dataQuantity := 1000

	file, err := os.Create(startTime.String())
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// generating records

	err = writer.Write(header)
	if err != nil {
		fmt.Printf("Error writing header: %v\n", err)
		return
	}

	var wg sync.WaitGroup

	for i := 0; i < dataQuantity; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			fmt.Printf("Generating record %d...\n", i+1)
			err = writer.Write([]string{
				nameSet[0],
				fmt.Sprintf("%d", i), // Age between 20 and 69
				stateSet[0],
			})

			if err != nil {
				fmt.Printf("Error writing record %d.\n", i)
			}
		}()
	}

	wg.Wait()

	elapsedTime := time.Since(startTime)
	fmt.Printf("Data generation completed in %v\n", elapsedTime.Minutes())
	fmt.Printf("File 'customers.csv' created successfully with %d records.\n", dataQuantity)

	// Implement the logic for generating data here
	// This function will be called when the user selects "Gerador de dados"
	//https://medium.com/@j.d.livni/create-a-load-bar-in-go-f158837ff4c4
}
