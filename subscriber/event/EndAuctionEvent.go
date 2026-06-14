package event

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type EndAuctionEvent struct {
	Winner    common.Address // indexed
	AuctionId *big.Int       // indexed
	Amount    *big.Int
}
