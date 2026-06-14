package vo

import (
	"MyNFT/dao/entity"
	"time"
)

type AuctionVO struct {
	Id          int64
	Seller      string
	NftContract string
	NftTokenId  int64
	StartPrice  float64
	StartTime   time.Time
	EndTime     time.Time
	Status      int8
	Winner      *string
	WinPrice    *float64
}

func NewAuctionVO(entity *entity.Auction) *AuctionVO {
	return &AuctionVO{
		Id:          entity.Id,
		Seller:      entity.Seller,
		NftContract: entity.NftContract,
		NftTokenId:  entity.NftTokenId,
		StartPrice:  entity.StartPrice,
		StartTime:   entity.StartTime,
		EndTime:     entity.EndTime,
		Status:      entity.Status,
		Winner:      entity.Winner,
		WinPrice:    entity.WinPrice,
	}
}
