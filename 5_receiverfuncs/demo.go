// Demo - Receiver functions or struct methods

package main

import "fmt"

type shopping struct{

    item string
    quantity int
    price float64
}

// Methods are customised behavior to structs and provide abstraction

func (s shopping) shoppingInfo(){
    fmt.Println(s.item, s.quantity, s.price)
}

func (b shopping) sumTotal(b2 shopping) float64{
    return b.price + b2.price
}

func (b *shopping) changeQuantity(n int){
    (*b).quantity = n
}

func main(){

    basket := shopping{
                item: "apple",
                quantity: 6,
                price: 150.7,
    }
    basket2 := shopping{
                item : "orange",
                quantity: 12,
                price: 98.9,
    }
    basket.shoppingInfo()

    total := basket.sumTotal(basket2)
    fmt.Println(total)

    total = basket2.sumTotal(basket)
    fmt.Println(total)

    basket.changeQuantity(8)
    fmt.Println(basket)
}