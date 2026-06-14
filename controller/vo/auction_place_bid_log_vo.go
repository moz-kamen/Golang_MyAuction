package vo

import (
	"MyNFT/dao/entity"
	"time"
)

type AuctionPlaceBidLogVO struct {
	Id        int64
	AuctionId int64
	Bidder    string
	BidPrice  float64
	BidTime   time.Time
}

func NewAuctionPlaceBidLogVO(entity *entity.AuctionPlaceBidLog) *AuctionPlaceBidLogVO {
	return &AuctionPlaceBidLogVO{
		Id:        entity.Id,
		AuctionId: entity.AuctionId,
		Bidder:    entity.Bidder,
		BidPrice:  entity.BidPrice,
		BidTime:   entity.BidTime,
	}
}
