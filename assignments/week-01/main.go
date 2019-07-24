package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		text := scanner.Text()
		//TODO
		err := eval(text)
		if err != nil {
			// log.Fatal(err)
			fmt.Println("ERROR:", err)
		}
		fmt.Print("> ")
	}
}
