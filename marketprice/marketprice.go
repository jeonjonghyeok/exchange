package main

import (
	. "fmt"

	"time"

	log "github.com/sirupsen/logrus"
	gecko "github.com/superoo7/go-gecko/v3"
)

func GetMarketPrice() {

	sp, err := cg.SimplePrice(ids, vc)
	if err != nil {
		log.Fatal("not get coin price")
	}
	bitcoin := (*sp)["bitcoin"]
	eth := (*sp)["ethereum"]
	log.Info(Sprintf("Bitcoin is worth %f usd (krw %f)", bitcoin["usd"], bitcoin["krw"]))
	log.Info(Sprintf("Ethereum is worth %f usd (krw %f)", eth["usd"], eth["krw"]))

}

var cg *gecko.Client
var ids []string
var vc []string

func main() {
	cg = gecko.NewClient(nil)

	ids = []string{"bitcoin", "ethereum"}
	vc = []string{"usd", "krw"}

	for {
		go GetMarketPrice()
		time.Sleep(time.Second)
	}

}
