package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Ferret struct {
	Name string
	Food Food
}

type Food struct {
	Name   string
	Volume int
}

func main() {
	var ferrets []Ferret

	bytes, err := ioutil.ReadFile("ferret.json")
	if err != nil {
		fmt.Println(err)
	}

	errJson := json.Unmarshal(bytes, &ferrets)
	if errJson != nil {
		fmt.Println(errJson)
	}

	for _, value := range ferrets {
		fmt.Printf("%s\t%s\t%d\n", value.Name, value.Food.Name, value.Food.Volume)
	}
}

/* ferret.json
[
  {
    "Name": "maru",
    "Food": {
      "Name": "zupurimeGrainFreeDiet",
      "Volume": 5
    }
  },
  {
    "Name": "hana",
    "Food": {
      "Name": "ferretSelectionSenior",
      "Volume": 3
    }
  },
  {
    "Name": "natu",
    "Food": {
      "Name": "totalyComplete",
      "Volume": 7
    }
  }
]
*/


/* 実行結果
maru	zupurimeGrainFreeDiet	5
hana	ferretSelectionSenior	3
natu	totalyComplete	7
*/
