package main

import (
	"log"
	"net"
	. "fmt"
	"github.com/tutorialedge/go-grpc-beginners-tutorial/chat"
	"google.golang.org/grpc"
	"quote"
)
type Wallet struct {
	crypto map[string]float32
	krw float32
}
type Order struct {
	status string
	isbuy string
	amount float32
	price float32
	ch int
}
func newWallet() *Wallet {
	w:= Wallet{}
	w.krw = 100000
	w.crypto = map[string]float32{}
	w.crypto["bitcoin"]=10000
	return &w
}
func orderTx(choiceCrypto int, amount float32, price float32, chs []chan float32,heap []Order, order string,w *Wallet) {
	defer func() {
		if r:=recover(); r!=nil {
			Println(r)
		}
	}()
	var crypto string
	var ch, nch int
	switch choiceCrypto {
		case 1:
			crypto = "bitcoin"
		case 2:
			crypto = "ethereum"
		case 3:
			crypto = "tether"
		case 4:
			crypto = "ripple"
		case 5:
			crypto = "polkadot"
		case 6:
			crypto = "litecoin"
		case 7:
			crypto = "bitcoin-cash"
	}
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
				w.crypto[crypto]+=amount
				w.krw-=price*amount
			}else if order=="매도" {
				w.crypto[crypto]-=amount
				w.krw+=price*amount
			}
			Println(order,"주문이 채결되었습니다.")
			break
		}else if chdata == amount {
			chs[ch]<-0
			heap[ch] = Order{}
			if order=="매수" {
				w.crypto[crypto]+=amount
				w.krw-=price*amount
			}else if order=="매도" {
				w.crypto[crypto]-=amount
				w.krw+=price*amount
			}
			Println(order,"주문이 체결되었습니다.")
			break
		}else if chdata < amount {
			amount -= chdata
			chs[ch]<-0
			heap[ch].status="주문중"
			heap[ch].isbuy=order
			heap[ch].amount = amount
			heap[ch].price = price
			heap[ch].ch = ch
			chs[ch]<-amount
		} else if chdata > amount {
			chs[ch]<-amount
			if order=="매수" {
				w.crypto[crypto]+=amount
				w.krw-=price*amount
			}else if order=="매도" {
				w.crypto[crypto]-=amount
				w.krw+=price*amount
			}
			break
		}
	}
}
func order(choiceCrypto int, chs []chan float32,heap []Order,w *Wallet) {
	defer func() {
		if r:= recover(); r!=nil {
			Println(r)
		}
	}()

	var amount, price float32
	var order string
	var crypto string
	switch choiceCrypto {
		case 1:
			crypto = "bitcoin"
		case 2:
			crypto = "ethereum"
		case 3:
			crypto = "tether"
		case 4:
			crypto = "ripple"
		case 5:
			crypto = "polkadot"
		case 6:
			crypto = "litecoin"
		case 7:
	}
	Println()
	Println("타입 | 구매량 | 가격을 순서대로 입력하세요.")
	Println("ex) 매도 100 200")
	Scanln(&order,&amount,&price)
	if order == "매도" || order == "매수"{
		if order == "매도" && w.crypto[crypto]<amount	{
			panic("보유 화폐 부족")

		}
		if order == "매수" && w.krw < price*amount {
			panic("원화가 부족합니다.")
		}

		Println("보유원화:",w.krw)
		Println("총 금액:",price*amount)
		go orderTx(choiceCrypto,amount,price,chs,heap, order,w)
	} else {
		panic("잘못입력하였습니다.")
	}

}
func wallet(w *Wallet) {
	Println("지갑 확인")
	//w.crypto["bitcoin"] = 1.2221
	Println("원화 보유량:",w.krw)
	for key,val:=range w.crypto {
		Println(key,"보유량:",val)
	}
}
func checkOrder(heap [][]Order, cryptosname []string) {
	for i:=0;i<len(cryptosname);i++ {
				Printf("------%s-----\n",cryptosname[i])
		for j:=0;j<len(heap[0]);j++ {
			if heap[i][j].status == "주문중" {
				Println(heap[i][j])
			}
		}
		Println()
	}
}
func orderMenu(cryptos []quote.Crypto, chs [][]chan float32,heap [][]Order,w *Wallet, cryptosname []string) {
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
			order(choiceCrypto,chs[choiceCrypto-1],heap[choiceCrypto-1],w)
}
func main() {
	defer func() {
		if r:= recover(); r!=nil {
			Println(r)
		}
	}()
	w := newWallet()


	var chs = make([][]chan float32,7)
	var heap = make([][]Order,7)


	for i:=0;i<7;i++ {
		chs[i] = make([]chan float32, 10)
		heap[i] = make([]Order,10)
		for k:=0;k<10;k++ {
			chs[i][k] = make(chan float32)
		}
	}

		var cryptosname []string = []string{"bitcoin","ethereum","tether","ripple","polkadot","litecoin","bitcoin-cash"}
	for {

	Println("2. 지갑 확인")
	Println("3. 주문")
	Println("4. 주문 확인")
	Println("5. 프로그램 종료")

	var num int
	Scanln(&num)

	switch num {
		case 2:
			Println("**********지갑***********")
			wallet(w)
			Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			Scanln()
		case 3:
			Println("**********주문***********")
			orderMenu(quote.GetCryptos(),chs,heap,w,cryptosname)
			Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			Scanln()
		case 4:
			Println("주문확인")
			checkOrder(heap,cryptosname)
			Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			Scanln()
		case 5:
			panic("프로그램종료")
			Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			Scanln()
		default:
			Println("잘못 입력하였습니다.")
			Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			Scanln()

	}
}
	lis, err := net.Listen("tcp",Sprintf(":%d",9010))
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
