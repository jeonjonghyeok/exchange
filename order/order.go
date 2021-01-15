package main
import (
	"order"
	. "fmt"
)
func main() {
	var isbuy string
	var price, amount float32
for {
	Println("가격 구매량, 구분")
	Println("ex) 3000 2 buy")
	Scanln(&price,&amount,&isbuy)
		o:= order.NewOrder(price,amount,isbuy)
		Println(o)
		Println("총 금액:",price*amount)
		Println("계속 주문하려면 enter")
		Scanln()
}
}

