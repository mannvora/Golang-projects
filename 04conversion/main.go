package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Please enter the rating: ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("The rating is: ", input)
}
