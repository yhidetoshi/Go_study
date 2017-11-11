package main

import "fmt"

const N = 24

func calc_time(in_n int)int{

	var result int

	if in_n <= N && in_n >= 0{
		result = in_n
	}else{	
		result = in_n - N
	}
		return result
}


func main(){
	var in_n int
	var valu int

	fmt.Scan(&in_n)

	valu = calc_time(in_n)
	fmt.Println(valu)
}

