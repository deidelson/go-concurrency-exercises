package main

import "fmt"

const channel_size = 10

//En los canales con buffer no es necesario que exista una goroutine "escuchando" el canal
func main() {
	c := make(chan int, channel_size)

	populateBufferedChannel(c)

	for item := range c {
		fmt.Println(item)
	}
}

func populateBufferedChannel(c chan <- int  ) {
	defer close(c)
	for i := 0; i<channel_size; i++ {
		c <- i
	}
}
