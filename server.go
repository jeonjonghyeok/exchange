package main

import (
	"log"
	"net"
	. "fmt"
	"github.com/tutorialedge/go-grpc-beginners-tutorial/chat"
	"google.golang.org/grpc"
	gecko "github.com/superoo7/go-gecko/v3"
)

type crypto struct{
	name string
	currentprice float32
	price map[string]int
}
type wallet struct {
	krw float32
	crypto map[string]float32
}

func quote() {
	choiceCrypto:=0
	for {
		Println("1. 비트코인: ")
		Println("2. 이더리움: ")

		cg:= gecko.NewClient(nil)

		ids:= []string{"bitcoin","ethereum"}
		vc:= []string{"usd","myr"}
		sp, err := cg.SimplePrice(ids,vc)
		if err!=nil {
			log.Fatal(err)
		}
		bitcoin := (*sp)["bitcoin"]
		eth := (*sp)["ethereum"]
		Println(Sprintf("Bitcoin is worth %f usd (myr %f)", bitcoin["usd"],bitcoin["myr"]))
		Println(Sprintf("Ethereum is worth %f usd (myr %f)", eth["usd"],eth["myr"]))


		Println("거래할 화폐 선택")
		Scanln(&choiceCrypto)
		switch choiceCrypto {
			case 1:
				order(1)
			case 2:
				order(2)
			default :
				break
		}
	}

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
			quote()
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
			Println("**********프로그램 종료***********")
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
