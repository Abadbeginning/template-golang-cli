package main

import (
	"fmt"
)

func goroutineDemo() {
	go fmt.Println("2")
	fmt.Println("1")
	//time.Sleep(time.Second)
}

func channelDemo() {
	c := make(chan string)

	go func() {
		fmt.Println("2")
		c <- "发送"
	}()

	fmt.Println("1")

	v := <-c
	fmt.Println("接收值：", v)
}

func main() {
	//goroutineDemo()
	channelDemo()
}
