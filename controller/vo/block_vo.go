package vo

import (
	"time"

	"github.com/ethereum/go-ethereum/core/types"
)

type BlockVO struct {
	BlockNumber uint64
	Hash        string
	ParentHash  string
	Timestamp   time.Time
	TxCount     int
	GasUsed     uint64
	GasLimit    uint64
}

func NewBlockVO(block *types.Block) *BlockVO {
	return &BlockVO{
		BlockNumber: block.NumberU64(),
		Hash:        block.Hash().Hex(),
		ParentHash:  block.ParentHash().Hex(),
		Timestamp:   time.Unix(int64(block.Time()), 0).UTC(), // 转为 UTC 时间
		TxCount:     len(block.Transactions()),
		GasUsed:     block.GasUsed(),
		GasLimit:    block.GasLimit(),
	}
}
