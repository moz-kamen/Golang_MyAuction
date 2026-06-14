package vo

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type TransactionVO struct {
	Hash      string
	From      string
	To        string
	Value     string
	Gas       uint64
	InputData string
	Status    uint64
	GasUsed   uint64
	LogCount  int
}

func NewSimpleTransactionVO(transaction *types.Transaction, from common.Address) *TransactionVO {
	toAddress := ""
	if transaction.To() != nil {
		toAddress = transaction.To().Hex()
	}

	return &TransactionVO{
		Hash:  transaction.Hash().Hex(),
		From:  from.Hex(),
		To:    toAddress,
		Value: transaction.Value().String(),
		Gas:   transaction.Gas(),
	}
}

func NewTransactionVO(transaction *types.Transaction, from common.Address, receipt *types.Receipt) *TransactionVO {
	toAddress := ""
	if transaction.To() != nil {
		toAddress = transaction.To().Hex()
	}

	return &TransactionVO{
		Hash:      transaction.Hash().Hex(),
		From:      from.Hex(),
		To:        toAddress,
		Value:     transaction.Value().String(),
		Gas:       transaction.Gas(),
		InputData: hex.EncodeToString(transaction.Data()),
		Status:    receipt.Status,
		GasUsed:   receipt.GasUsed,
		LogCount:  len(receipt.Logs),
	}
}
