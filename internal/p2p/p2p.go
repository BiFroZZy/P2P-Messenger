package p2p

import (
	"context"
	
	libp2p "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/p2p/security/noise"
	"github.com/libp2p/go-libp2p/core/peer"
	_"github.com/libp2p/go-libp2p/core/network"

	l "messenger/internal/logger"
)

var logger = l.InitLogger()

func CreateSourceNode() host.Host{
	node, err := libp2p.New()
	if err != nil{
		logger.Error().Err(err).Msg("Error occured while creating source node")
	}
	logger.Info().Msg("Host created!"+node.ID().String())
	return node
}

func CreateTargetNode() host.Host{
	node, err := libp2p.New(libp2p.Security(noise.ID, noise.New))
	if err != nil{
		logger.Error().Err(err).Msg("Error occurde while creating target node")
	}
	logger.Info().Msg("Host created!"+node.ID().String())
	return node
}

func ConnectingToPeers(host host.Host){
	peerAddr, err := peer.AddrInfoFromString("/ip4/127.0.0.1/tcp/9000/p2p/")
	if err != nil{
		logger.Error().Err(err).Msg("Error occured while getting peers")
	}
	if err = host.Connect(context.Background(), *peerAddr); err != nil {
		logger.Error().Err(err).Msg("Error occured while connecting to peers")
	}
}

func NewStream(host host.Host, peerID peer.ID, message string){
	stream, err := host.NewStream(context.Background(), peerID, "/my-protocol/1.0.0")
	if err != nil {
		logger.Error().Err(err).Msg("Error occured while creating stream")
	}

	_, err = stream.Write([]byte(message))
	if err != nil {
		logger.Error().Err(err).Msg("Failed to write message")
	}
}

func GetMessage(){
	
}