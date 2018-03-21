package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./inputData.csv")
	if err != nil {
		log.Fatal("Error", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		} else {
			if err != nil {
				log.Fatal("Error", err)
			}
		}
		fmt.Println(record)
	}

}
