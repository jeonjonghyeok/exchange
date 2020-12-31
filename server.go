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

type crypto struct{
	name string
	price float32
}
type wallet struct {
	krw float32
	crypto map[string]float32
}
func quoteGet(client *redis.Client, cryptos []crypto) {
		cg:= gecko.NewClient(nil)
		ids:= []string{"bitcoin","ethereum"}
		vc:= []string{"usd","myr"}
		for {
			time.Sleep(time.Second*1)
		sp, err := cg.SimplePrice(ids,vc)
		if err!=nil {
			log.Fatal(err)
		}
		for i:=0;i<2;i++ {
			cryptos[i].price = (*sp)[cryptos[i].name]["usd"]
			err = client.Set(client.Context(),Sprint("%s",cryptos[i].name),cryptos[i].price,0).Err()
		}
		if err != nil {
			Println(err)
		}
	}

}

func quote(client *redis.Client,cryptos []crypto) {


		for i:=0;i<2;i++ {
		val, err := client.Get(client.Context(),Sprint("%s",cryptos[i].name)).Result()
		Println(cryptos[i].name,val)
			if err != nil {
				Println(err)
			}
		}

		Scanln()

}
func order(choiceCrypto int) {
	Println(choiceCrypto,"를 선택하셨습니다.")
	Println("1. 매도")
	Println("2. 매수")
	Scanln()
}
func sellCrypto() {
	Println("매도")
}
func buyCrypto() {
	Println("매수")
}
func main() {
	defer func() {
		if r:= recover(); r!=nil {
			Println(r)
		}
	}()
/*	client:= redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	*/

	//pong, err := client.Ping(client.Context()).Result()
	//Println(pong,err)
	client:= redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
		cryptos := make([]crypto,5)
		cryptos[0].name = "bitcoin"
		cryptos[1].name = "ethereum"
		go quoteGet(client, cryptos)
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
			quote(client,cryptos)
			Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			Scanln()
		case 2:
			Println("**********지갑***********")
			Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			Scanln()
		case 3:
			Println("**********매도 주문***********")
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
