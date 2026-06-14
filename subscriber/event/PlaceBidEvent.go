package event

import (
	"MyNFT/dao/entity"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type PlaceBidEvent struct {
	Bidder    common.Address // indexed
	AuctionId *big.Int       // indexed
	Amount    *big.Int
}

func (event *PlaceBidEvent) ParseEntity() entity.AuctionPlaceBidLog {
	var auctionLog = entity.AuctionPlaceBidLog{}
	auctionLog.AuctionId = event.AuctionId.Int64()
	auctionLog.Bidder = event.Bidder.Hex()
	auctionLog.BidPrice, _ = event.Amount.Float64()
	auctionLog.BidTime = time.Now()
	return auctionLog
}
