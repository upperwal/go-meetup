package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/savaki/jq"
	"github.com/upperwal/go-meetup/p2p"
	"github.com/upperwal/go-meetup/structs"
	"github.com/upperwal/go-meetup/utils"
)

var op, _ = jq.Parse(".")
var p2pNet *p2p.P2P

func main() {
	/* logging.SetLogLevel("svc-bootstrap", "DEBUG")
	logging.SetLogLevel("application", "DEBUG") */

	port := flag.String("p", "3000", "port to start the server")
	flag.Parse()

	ctx := context.Background()

	var err error
	p2pNet, err = p2p.NewP2P(ctx)
	if err != nil {
		panic(err)
	}

	/* b := structs.Block{
		Nonce:     1,
		Timestamp: time.Now().String(),
		Transactions: []structs.Transaction{
			structs.Transaction{
				Value: 1,
			},
		},
	} */

	r := mux.NewRouter()
	r.HandleFunc("/", entireBlockchain)
	r.HandleFunc("/send_transaction", sendTrans)
	r.HandleFunc("/state", readState)

	log.Fatal(http.ListenAndServe(":"+(*port), r))

	p2pNet.BCService.SendTransaction(structs.Transaction{
		Value: -1,
	})
	p2pNet.BCService.SendTransaction(structs.Transaction{
		Value: 5,
	})
}

func entireBlockchain(w http.ResponseWriter, r *http.Request) {
	fmt.Println(utils.GetPrettyJSON(p2pNet.BCService.BlockchainStore.Blocks))
	m, _ := p2pNet.BCService.MarshalBlockchain()

	w.Header().Set("Content-Type", "application/json")
	w.Write(m)
}

func readState(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(utils.GetPrettyJSON(p2pNet.BCService.BlockchainStore.Blocks))
	s := p2pNet.BCService.BlockchainStore.GetState()

	state := `{"State":` + s + `}`

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(state))
}

func sendTrans(w http.ResponseWriter, r *http.Request) {
	var t structs.Transaction
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		panic(err)
	}
	p2pNet.BCService.SendTransaction(t)

	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(t)
	w.Write(d)
}
