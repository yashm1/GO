package main

import "fmt"

func main() {
	msgs := make(chan string, 2)
	msgs <- "hello"
	msgs <- "Bye"
	fmt.Println(<-msgs)
	msgs <- "abc"
	fmt.Println(<-msgs)
}
