package main

import (
	"fmt"
)

//En los canales sin buffer es necesario que exista una goroutine escuchando el canal, en caso contrario Panic
func main() {
	c := populateChannelAsyncReturning()

	for value := range c {
		fmt.Println(value)
	}


}

//3 opciones para llenar canales sin buffer, en los 3 casos el receptor debe estar listo (puede ser en el mismo main)


//1 Creear una funcion normal y llamarla desde afuera de forma asincronica: go populateChannel(c chan int)
func populateChannel(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

//2- Recibir el canal y llenarlo de forma asincronica: populateChannelAsync(c)
func populateChannelAsync(c chan int) {
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
}

//3- Crear el canal, llenarlo asincronicamente y devolverlo: c := populateChannelAsyncReturning()
func populateChannelAsyncReturning() chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

