module github.com/upperwal/go-meetup

go 1.12

require (
	github.com/TylerBrock/colorjson v0.0.0-20180527164720-95ec53f28296
	github.com/ethereum/go-ethereum v1.9.2
	github.com/fatih/color v1.7.0 // indirect
	github.com/gorilla/mux v1.7.3
	github.com/ipfs/go-log v0.0.1
	github.com/libp2p/go-libp2p v0.3.1
	github.com/libp2p/go-libp2p-core v0.2.2
	github.com/pravahio/go-mesh v0.0.5
	github.com/savaki/jq v0.0.0-20161209013833-0e6baecebbf8
)

replace github.com/libp2p/go-libp2p-pubsub v0.1.0 => github.com/upperwal/go-libp2p-pubsub v0.1.1-0.20190822125434-affd4e4c6c42
