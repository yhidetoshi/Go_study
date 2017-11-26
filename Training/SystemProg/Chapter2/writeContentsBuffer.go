package Chapter2

import (
	"fmt"
	"bytes"
)

func main(){
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example"))
	fmt.Println(buffer.String())
}
