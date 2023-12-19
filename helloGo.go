package main
import (
	"fmt"
	"modules.com/packs"
)

func main(){
	message := packs.Hello("Hello Golang!!!")
	fmt.Printf(message)
}