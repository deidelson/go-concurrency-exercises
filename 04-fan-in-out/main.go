package main

import (
	"fmt"
	"sync"
)

type item struct {
	name string
	price float64
	discount bool
}

func main() {
	itemChannel := createChannelFromItems(createItems())

	//fanOut dividir un canal en varios para procesar
	discountChannel1 := applyDiscount(itemChannel)
	discountChannel2 := applyDiscount(itemChannel)

	//fanIn mergear la informacion de muchos canales en uno
	merge := fanIn(discountChannel1, discountChannel2)

	for element := range merge {
		fmt.Println(element)
	}
}

func createChannelFromItems(items []item) <- chan item  {
	out := make(chan item)

	go func() {
		defer close(out)
		for _, item := range items {
			out <- item
		}
	}()

	return out
}

func createItems() []item {
	items := make([]item, 0, 0)

	items = append(items, item{
		name: "tshirt",
		price: 20,
		discount: false,
	})
	items = append(items, item{
		name: "trausers",
		price: 30,
		discount: true,
	})
	items = append(items, item{
		name: "shoes",
		price: 100,
		discount: true,
	})

	return items
}

//Fan out function
func applyDiscount(itemChannel <-chan item) <- chan item {
	out := make(chan item)

	go func() {
		defer close(out)
		for value := range itemChannel {
			if value.discount {
				value.price = value.price * 0.5
			}
			out <- value
		}
	}()
	return out
}

//Recibe una lista de canales y hace un merge
func fanIn(channels ...<-chan item) <-chan item{
	var wg sync.WaitGroup
	out := make(chan item)

	wg.Add(len(channels))

	for _, channel := range channels {
		fillWithChannelData(&wg, out, channel)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func fillWithChannelData(wg *sync.WaitGroup, out chan <- item, filler <-chan item) {
	go func() {
		defer wg.Done()
		for element := range filler {
			out <- element
		}
	}()
}




