package Chapter2

import (
	"os"
)

func main(){
	file, err := os.Create("test.txt")
	if err != nil{
		panic(err)
	}
	file.Write([]byte("os.FIle example"))
	file.Close()
}