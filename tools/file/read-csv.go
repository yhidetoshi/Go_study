package main

import(
	"encoding/csv"
	"io"
	"log"
	"os"
	"fmt"
)

func main(){
	file, err := os.Open("./sample-csv.csv")

	if err != nil{
		log.Fatal("Error", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}else{
			if err != nil{
				log.Fatal("Error", err)
			}
		}

		//log.Printf("%#v", record)
		fmt.Println(record)

	}
}
