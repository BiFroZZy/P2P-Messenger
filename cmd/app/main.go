package main

import (
	"log"
	"context"

	libp2p "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/p2p/security/noise"
	"github.com/libp2p/go-libp2p/core/peer"
	_"github.com/libp2p/go-libp2p/core/network"
)

func main(){
	host, err := libp2p.New(
		libp2p.Security(noise.ID, noise.New),
	)
	if err != nil{
		log.Printf("Can't create host: %v", err)
	} 
	defer host.Close()
	
	log.Println("Host created!", host.ID())

	peerAddr, err := peer.AddrInfoFromString("/ip4/127.0.0.1/tcp/9000/p2p/")
	if err != nil{
		log.Printf("Can't connect to peer: %v", err)
	}
	if err = host.Connect(context.Background(), *peerAddr); err != nil {
		log.Fatalf("Ошибка подключения: %v", err)
	}
	

	// stream, err := host.NewStream(context.Background(), peerID, "/my-protocol/1.0.0")
	// if err != nil {
	// 	log.Fatalf("Ошибка создания потока: %v", err)
	// }

	// _, err = stream.Write([]byte("Привет, P2P!"))
	// if err != nil {
	// 	log.Fatalf("Ошибка записи: %v", err)
	// }


	
}
