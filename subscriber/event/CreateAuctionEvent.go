package event

import (
	"MyNFT/dao/entity"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type CreateAuctionEvent struct {
	Seller        common.Address // indexed
	AuctionId     *big.Int       // indexed
	NftContract   common.Address
	TokenId       *big.Int
	StartPrice    *big.Int
	DurationHours *big.Int
}

func (event *CreateAuctionEvent) ParseEntity() entity.Auction {
	var auction = entity.Auction{}
	auction.Id = event.AuctionId.Int64()
	auction.Seller = event.Seller.Hex()
	auction.NftContract = event.NftContract.Hex()
	auction.NftTokenId = event.TokenId.Int64()
	auction.StartPrice, _ = event.StartPrice.Float64()
	auction.StartTime = time.Now()
	auction.EndTime = auction.StartTime.Add(time.Duration(event.DurationHours.Int64()) * time.Hour)
	return auction
}
