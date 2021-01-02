package main

import (
	"log"
	"net"
	. "fmt"
	"github.com/tutorialedge/go-grpc-beginners-tutorial/chat"
	"google.golang.org/grpc"
	gecko "github.com/superoo7/go-gecko/v3"
	"github.com/go-redis/redis"
	//"time"
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
	amount float32
	price float32
	ch int
}
// heap
type minheap struct {
    heapArray []Order
    size      int
    maxsize   int
}

func newMinHeap(maxsize int) *minheap {
    minheap := &minheap{
        heapArray: []Order{},
        size:      0,
        maxsize:   maxsize,
    }
    return minheap
}

func (m *minheap) leaf(index int) bool {
    if index >= (m.size/2) && index <= m.size {
        return true
    }
    return false
}

func (m *minheap) parent(index int) int {
    return (index - 1) / 2
}

func (m *minheap) leftchild(index int) int {
    return 2*index + 1
}

func (m *minheap) rightchild(index int) int {
    return 2*index + 2
}

func (m *minheap) insert(item Order) error {
    if m.size >= m.maxsize {
        return Errorf("Heap is ful")
    }
    m.heapArray = append(m.heapArray, item)
    m.size++
    m.upHeapify(m.size - 1)
    return nil
}

func (m *minheap) swap(first, second int) {
    temp := m.heapArray[first]
    m.heapArray[first] = m.heapArray[second]
    m.heapArray[second] = temp
}

func (m *minheap) upHeapify(index int) {
    for m.heapArray[index].price < m.heapArray[m.parent(index)].price {
        m.swap(index, m.parent(index))
        index = m.parent(index)
    }
}

func (m *minheap) downHeapify(current int) {
    if m.leaf(current) {
        return
    }
    smallest := current
    leftChildIndex := m.leftchild(current)
    rightRightIndex := m.rightchild(current)
    //If current is smallest then return
    if leftChildIndex < m.size && m.heapArray[leftChildIndex].price < m.heapArray[smallest].price {
        smallest = leftChildIndex
			}
    if rightRightIndex < m.size && m.heapArray[rightRightIndex].price < m.heapArray[smallest].price {
        smallest = rightRightIndex
    }
    if smallest != current {
        m.swap(current, smallest)
        m.downHeapify(smallest)
    }
    return
}
func (m *minheap) buildMinHeap() {
    for index := ((m.size / 2) - 1); index >= 0; index-- {
        m.downHeapify(index)
    }
}

func (m *minheap) remove() Order {
    top := m.heapArray[0]
    m.heapArray[0] = m.heapArray[m.size-1]
    m.heapArray = m.heapArray[:(m.size)-1]
    m.size--
    m.downHeapify(0)
    return top
}



func newWallet() *Wallet {
	w:= Wallet{}
	w.krw = 100000
	w.crypto = map[string]float32{}
	w.crypto["bitcoin"]=10000
	return &w
}
func setQuote(client *redis.Client, cryptos []Crypto,ids []string) {
		cg:= gecko.NewClient(nil)
		//ids:= []string{"bitcoin","ethereum","tether","ripple","polkadot","litecoin","bitcoin-cash"}
		vc:= []string{"usd","krw"}

		for {
			//time.Sleep(time.Second*1)
		sp, err := cg.SimplePrice(ids,vc)
		if err!=nil {
			log.Fatal(err)
		}
		for i:=0;i<len(cryptos);i++ {
			cryptos[i].price = (*sp)[cryptos[i].name]["usd"]
			/*json, err1 := json.Marshal(cryptos[i])
			if err1 != nil {
				Println(err1)
			}
			*/
			err = client.Set(client.Context(),Sprint("%s",cryptos[i].name),cryptos[i].price,0).Err()
		}
		if err != nil {
			Println(err)
		}
	}
}
func getQuote(client *redis.Client,cryptos []Crypto) {
		for i:=0;i<len(cryptos);i++ {
		val, err := client.Get(client.Context(),Sprint("%s",cryptos[i].name)).Result()
		Println(i+1,cryptos[i].name,val)
			if err != nil {
				Println(err)
			}
		}
}
func sellOrder(choiceCrypto int, amount float32, price float32, chs []chan float32,heap []Order, order string,w *Wallet) {
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
	}
	ch = -1
	for i:=4;i>=0;i-- {
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
			crypto="bitcoin"
		case 2:
			crypto="ethereum"
	}

	Println("타입 구매량 가격")
	Println("ex) 매도 200 300")
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
		go sellOrder(choiceCrypto,amount,price,chs,heap, order,w)
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
func heapInsert(heap []Order,item Order) {



}
func heapDelete(heap []Order,item int) {


}
func checkOrder(heap []Order) {
	for i:=0;i<5;i++ {
		Println(heap[i])
	}

}
func main() {
	defer func() {
		if r:= recover(); r!=nil {
			Println(r)
		}
	}()
	w := newWallet()

	var bchs = make([]chan float32, 5)
	var echs = make([]chan float32, 5)
	bheap := make([]Order,5)
	eheap := make([]Order,5)
	for i:=0;i<5;i++ {
		bchs[i] = make(chan float32)
		echs[i] = make(chan float32)
	}

	client:= redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
		cryptos := make([]Crypto,7)
		var cryptosname []string = []string{"bitcoin","ethereum","tether","ripple","polkadot","litecoin","bitcoin-cash"}
		for j:=0;j<len(cryptos);j++ {
			cryptos[j].name = cryptosname[j]
		}

		go setQuote(client, cryptos,cryptosname)

	for {

	Println("1. 시세 확인")
	Println("2. 지갑 확인")
	Println("3. 주문")
	Println("4. 주문 확인")
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
			Println("**********주문***********")
			Println("구매할 Crypto 선택")
			choiceCrypto := 0
			getQuote(client,cryptos)
			Scanln(&choiceCrypto)
			if choiceCrypto==1 {
				order(choiceCrypto,bchs,bheap,w)
			}else if choiceCrypto==2 {
				order(choiceCrypto,echs,eheap,w)
			} else {
				panic("잘못입력하였습니다.")
			}
			Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			Scanln()
		case 4:
			Println("주문확인")
			Println("------비트코인-----")
			checkOrder(bheap)
			Println()
			Println("------이더리움-----")
			checkOrder(eheap)
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
