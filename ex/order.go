package main

import (
	"log"
	"net"
	. "fmt"
	"github.com/tutorialedge/go-grpc-beginners-tutorial/chat"
	"google.golang.org/grpc"
	"quote"
)
type Order struct {
	status string
	isbuy string
	amount float32
	price float32
	ch int
}
func orderTx(amount float32, price float32, chs []chan float32,heap []Order, order string) {
	defer func() {
		if r:=recover(); r!=nil {
			Println(r)
		}
	}()
	var ch, nch int
	ch = -1
	for i:=len(heap)-1;i>=0;i-- {
		if heap[i].status == "주문중" {
			if heap[i].price ==price &&heap[i].isbuy!=order{
				ch = heap[i].ch
			}
		}else {
			nch= i
		}
	}
	if ch== -1 {
		ch = nch
		heap[ch].status="주문중"
		heap[ch].isbuy=order
		heap[ch].amount=amount
		heap[ch].price=price
		heap[ch].ch = ch
		Println("주문이 접수되었습니다.")
		chs[ch]<-amount
	}
	for {
		chdata:=<-chs[ch]
		if chdata == 0 {
			heap[ch]=Order{}
			if order=="매수" {
			}else if order=="매도" {
			}
			Println(order,"주문이 채결되었습니다.")
			break
		}
	}
}
func order(choiceCrypto int, chs []chan float32,heap []Order) {
	defer func() {
		if r:= recover(); r!=nil {
			Println(r)
		}
	}()

	var amount, price float32
	var order string
	Println()
	Println("타입 | 구매량 | 가격을 순서대로 입력하세요.")
	Println("ex) 매도 100 200")
	Scanln(&order,&amount,&price)
	if order == "매도" || order == "매수"{
		go orderTx(amount,price,chs,heap, order)
	} else {
		panic("잘못입력하였습니다.")
	}
}
func orderMenu(cryptos []quote.Crypto, chs [][]chan float32,heap [][]Order, cryptosname []string) {
	defer func() {
		if r:= recover(); r!=nil {
			Println(r)
		}
	}()
			Println("구매할 Crypto 선택")
			choiceCrypto := 0
			quote.GetQuote(cryptosname)
			Scanln(&choiceCrypto)
			if choiceCrypto < 0 || choiceCrypto > 10 {
				panic("없는 코인입니다.")
			}
			order(choiceCrypto,chs[choiceCrypto-1],heap[choiceCrypto-1])
}
func main() {
	defer func() {
		if r:= recover(); r!=nil {
			Println(r)
		}
	}()

	var chs = make([][]chan float32,7)
	var heap = make([][]Order,7)
	var cryptosname []string = []string{"bitcoin","ethereum","tether","ripple","polkadot","litecoin","bitcoin-cash"}


	for i:=0;i<7;i++ {
		chs[i] = make([]chan float32, 10)
		heap[i] = make([]Order,10)
		for k:=0;k<10;k++ {
			chs[i][k] = make(chan float32)
		}
	}
	for {
		orderMenu(quote.GetCryptos(),chs,heap,cryptosname)
	}
	lis, err := net.Listen("tcp",Sprintf(":%d",9020))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s:= chat.Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer,&s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s",err)
	}
}
