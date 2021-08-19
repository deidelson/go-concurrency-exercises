//La idea es pasar información de una goroutine a otra usando canales
package main

import "fmt"

type item struct {
	name string
	price float64
	discount bool
}

func main() {

	itemChannel := make(chan item)

	go fillChannelWithItems(itemChannel, createItems())

	newItemChannel := applyDiscount(itemChannel)

	for processed := range newItemChannel {
		fmt.Println("item:", processed.name, "prince:", processed.price)
	}

}

//Caso 1 paso el channel desde afuera y lo lleno llamando a la funcion desde otra gorutine
func fillChannelWithItems(c chan <- item, items []item) {
	defer close(c)
	for _, value := range items {
		c <- value
	}
}

//Caso 2 retorno el channel y llamo a una goroutine desde adentro para llenarlo
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
