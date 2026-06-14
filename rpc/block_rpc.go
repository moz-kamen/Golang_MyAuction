package rpc

import (
	"MyNFT/config"
	"MyNFT/controller/vo"
	"MyNFT/util"
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type BlockRpc struct {
	Config *config.Config
}

func NewBlockRpc(config *config.Config) *BlockRpc {
	return &BlockRpc{
		Config: config,
	}
}

func (rpc *BlockRpc) BlockHash() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), rpc.Config.EthClient.Timeout*time.Second)
	defer cancel()
	client := util.BuildEthRpcClient(rpc.Config, ctx)

	blockNum, err := client.BlockNumber(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get block number %v: %w", blockNum, err)
	}

	block, err := client.BlockByNumber(ctx, new(big.Int).SetUint64(blockNum))
	if err != nil {
		return "", fmt.Errorf("failed to fetch block by number %v: %w", blockNum, err)
	}

	return block.Hash().Hex(), nil
}

func (rpc *BlockRpc) BlockByHash(blockHash string) (*vo.BlockVO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), rpc.Config.EthClient.Timeout*time.Second)
	defer cancel()
	client := util.BuildEthRpcClient(rpc.Config, ctx)

	block, err := client.BlockByHash(ctx, common.HexToHash(blockHash))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch block by hash %v: %w", blockHash, err)
	}
	return vo.NewBlockVO(block), nil
}
