package Chapter2

import(
	"os"
)

func main() {
	os.Stdout.Write([]byte("os.FIle example"))
}