package rpc

import (
	"MyNFT/config"
	"MyNFT/controller/vo"
	"MyNFT/util"
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type TransactionRpc struct {
	Config *config.Config
}

func NewTransactionRpc(config *config.Config) *TransactionRpc {
	return &TransactionRpc{
		Config: config,
	}
}

func (rpc *TransactionRpc) Transactions(blockHash string) (*[]vo.TransactionVO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), rpc.Config.EthClient.Timeout*time.Second)
	defer cancel()
	client := util.BuildEthRpcClient(rpc.Config, ctx)

	// 查询区块信息
	block, err := client.BlockByHash(ctx, common.HexToHash(blockHash))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch block by hash %v: %w", blockHash, err)
	}

	// 构造签名对象
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chainID: %v", err)
	}
	signer := types.NewLondonSigner(chainID) // 兼容现代 EIP-1559 交易

	// 构造交易记录
	var transactionVOList []vo.TransactionVO
	// 遍历交易并转换为VO
	for _, transaction := range block.Transactions() {
		// 解析 From 地址（需要通过签名反解）
		from, err := types.Sender(signer, transaction)
		if err != nil {
			return nil, fmt.Errorf("failed to get From: %v", err)
		}
		transactionVOList = append(transactionVOList, *vo.NewSimpleTransactionVO(transaction, from))
	}

	return &transactionVOList, nil
}

func (rpc *TransactionRpc) TransactionByHash(transactionHash string) (*vo.TransactionVO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), rpc.Config.EthClient.Timeout*time.Second)
	defer cancel()
	client := util.BuildEthRpcClient(rpc.Config, ctx)

	transaction, _, err := client.TransactionByHash(ctx, common.HexToHash(transactionHash))
	if err != nil {
		return nil, fmt.Errorf("tx not found or rpc error: %v", err)
	}

	// 构造签名对象
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chainID: %v", err)
	}
	signer := types.NewLondonSigner(chainID) // 兼容现代 EIP-1559 交易

	// 解析 From 地址（需要通过签名反解）
	from, err := types.Sender(signer, transaction)
	if err != nil {
		return nil, fmt.Errorf("failed to get From: %v", err)
	}

	// 查询交易回执
	receipt, err := client.TransactionReceipt(ctx, common.HexToHash(transactionHash))
	if err != nil {
		return nil, fmt.Errorf("查询交易回执失败: %w", err)
	}

	return vo.NewTransactionVO(transaction, from, receipt), nil
}
