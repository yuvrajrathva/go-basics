package main

import "fmt"

func main(){
	numbers := []string{"1", two()}
	numbers = append(numbers, "3")

	for i, n := range numbers{
		fmt.Println(i, n)
	}
}

func two() string{
	return "2"
}