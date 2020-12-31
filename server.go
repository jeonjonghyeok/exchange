package main

import (
	"log"
	"net"
	. "fmt"
	"github.com/tutorialedge/go-grpc-beginners-tutorial/chat"
	"google.golang.org/grpc"
	gecko "github.com/superoo7/go-gecko/v3"
	"github.com/go-redis/redis"
	"time"
)

type Crypto struct{
	name string
	price float32
}
type Wallet struct {
	crypto map[string]float32
	krw float32
}
type Order struct {
	status string
	isbuy string


}
func newWallet() *Wallet {
	w:= Wallet{}
	w.krw = 100000
	w.crypto = map[string]float32{}
	return &w
}
func setQuote(client *redis.Client, cryptos []Crypto) {
		cg:= gecko.NewClient(nil)
		ids:= []string{"bitcoin","ethereum"}
		vc:= []string{"usd","krw"}

		for {
			time.Sleep(time.Second*1)
		sp, err := cg.SimplePrice(ids,vc)
		if err!=nil {
			log.Fatal(err)
		}
		for i:=0;i<2;i++ {
			cryptos[i].price = (*sp)[cryptos[i].name]["usd"]
			/*json, err1 := json.Marshal(cryptos[i])
			if err1 != nil {
				Println(err1)
			}
			*/
			err = client.Set(client.Context(),Sprint("%s",cryptos[i].name),cryptos[i].price,0).Err()

			//err = client.Set(client.Context(),Sprint("%s",cryptos[i].name),cryptos[i].price,0).Err()
		}
		if err != nil {
			Println(err)
		}
	}

}

func getQuote(client *redis.Client,cryptos []Crypto) {


		for i:=0;i<2;i++ {
		val, err := client.Get(client.Context(),Sprint("%s",cryptos[i].name)).Result()
		Println(i+1,cryptos[i].name,val)
			if err != nil {
				Println(err)
			}
		}


}
func sellOrder() {


}
func order(choiceCrypto int) {

	var choicePrice, choiceAmount int
	Println("주문가 선택")
	Scanln(&choicePrice)
	Println("주문량 선택")
	Scanln(&choiceAmount)

	//go

}
func sellCrypto(client *redis.Client, cryptos []Crypto) {
	defer func() {
		if r:= recover(); r!=nil {
			Println(r)
		}
	}()
	Println("매도")
	getQuote(client,cryptos)
	var choiceCrypto int
	Println("매도할 Crypto를 선택하세요")
	Scanln(&choiceCrypto)

	switch choiceCrypto {
		case 1:
			Println("비트코인 주문")
		case 2:
			Println("이더리움 주문")
		default :
			panic("잘못 입력하였습니다.")

	}




}
func buyCrypto() {
	Println("매수")
}
func wallet(w *Wallet) {
	Println("지갑 확인")
	w.crypto["bitcoin"] = 1.2221
	for key,val:=range w.crypto {
		Println(key,"보유량:",val)
	}
}
func main() {
	defer func() {
		if r:= recover(); r!=nil {
			Println(r)
		}
	}()
	w := newWallet()

	client:= redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
		cryptos := make([]Crypto,5)
		cryptos[0].name = "bitcoin"
		cryptos[1].name = "ethereum"
		go setQuote(client, cryptos)
	for {

	Println("1. 시세 확인")
	Println("2. 지갑 확인")
	Println("3. 매도 주문")
	Println("4. 매수 주문")
	Println("5. 프로그램 종료")

	var num int
	Scanln(&num)

	switch num {
		case 1:
			Println("**********시세확인***********")
			getQuote(client,cryptos)
			Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			Scanln()
		case 2:
			Println("**********지갑***********")
			wallet(w)
			Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			Scanln()
		case 3:
			Println("**********매도 주문***********")
			sellCrypto(client,cryptos)
			Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			Scanln()
		case 4:
			Println("**********매수 주문***********")
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
