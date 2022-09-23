package main

import (
	"fmt"
	"time"
)

type Item struct {
	item_ID    int
	item_name  string
	item_price float32
}

func Product(qt int, pd_id int, pd_name string, pd_price float32) (id int, name string, price float32) {

	now := time.Now()

	getItem := Item{
		item_ID:    pd_id,
		item_name:  pd_name,
		item_price: pd_price}

	for i := 0; i < qt; i++ {
		if qt == 5 {
			sliItem := make([]int, id, 20)
			apItem := append(sliItem, 20, 55)
			fmt.Println(apItem, "Purchase of Date =>", now)
			break
		} else {
			id = getItem.item_ID + getItem.item_ID
			name = getItem.item_name
			price = getItem.item_price
			return
		}
	}
	return
}

func Bucket() {
	sp := Item{item_ID: 20}
	sp.item_name = "asus"
	fmt.Println(&sp)
}

func main() {
	/* fmt.Println(Product(5, 1, "apple", 25000.60)) */
	Bucket()

}
