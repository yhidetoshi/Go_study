package main

import (
	"sync"
	"fmt"
)

func main(){
	var wg sync.WaitGroup

	wg.Add(2)

	go func(){
		fmt.Println("Job1")
		wg.Done()
	}()
	go func(){
		fmt.Println("Job2")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Fin")
}

