package util

import (
	"MyNFT/config"
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func BuildEthRpcClient(config *config.Config, ctx context.Context) *ethclient.Client {
	client, err := ethclient.DialContext(ctx, config.EthClient.RpcUrl)
	if err != nil {
		log.Fatalf("failed to connect to Ethereum node: %v", err)
	}

	return client
}

func BuildEthWebsocketClient(config *config.Config, ctx context.Context) *ethclient.Client {
	client, err := ethclient.DialContext(ctx, config.EthClient.WebsocketUrl)
	if err != nil {
		log.Fatalf("failed to connect to Ethereum node: %v", err)
	}

	return client
}
