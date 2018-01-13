package main

import "fmt"

func main() {

	/* Long version */
	var languages map[string]string = map[string]string{"go": "golang", "rb": "ruby"}
	fmt.Println(languages)
	/* output
	map[go:golang rb:ruby]
	*/

	/* short version */
	languages := map[string]string{"go": "Golang", "rb": "ruby"}
	fmt.Println(languages)
	/* output
	map[go:golang rb:ruby]
	*/

	/* another 宣言*/
	languages := map[string]string{
		"go": "Golang",
		"rb": "ruby",
	}
	/* output
	map[go:golang rb:ruby]
	*/

	/* using make pattern */
	languages := make(map[string]string)
	languages["go"] = "Golang"
	languages["rb"] = "ruby"

}
