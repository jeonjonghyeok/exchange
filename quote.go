package main
import (
	"quote"
	. "fmt"
	"github.com/inancgumus/screen"
	"time"
)

func main() {

	var cryptosname []string = []string{"bitcoin","ethereum","tether","ripple","polkadot","litecoin","bitcoin-cash"}
	t := 0

	go quote.SetQuote(cryptosname)
	for {
		screen.Clear()
		Println("******현재 시세******")
		Println("경과시간:",t,"초")
		Println()
		quote.GetQuote(cryptosname)
		time.Sleep(time.Second*10)
		t=t+10
	}

	Scanln()



}
