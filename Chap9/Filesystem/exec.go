package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls")

	result, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// output ls
	fmt.Printf("%s\n", result)

}
