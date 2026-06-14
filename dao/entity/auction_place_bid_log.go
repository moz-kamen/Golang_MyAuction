package entity

import "time"

type AuctionPlaceBidLog struct {
	Id        int64     `gorm:"column:id;primaryKey;autoIncrement;comment:主键" json:"id"`
	AuctionId int64     `gorm:"column:auction_id;not null;index:idx_auction_id;comment:拍卖主键" json:"auction_id"`
	Bidder    string    `gorm:"column:bidder;not null;comment:竞拍者" json:"bidder"`
	BidPrice  float64   `gorm:"column:bid_price;not null;comment:竞拍价格" json:"bid_price"`
	BidTime   time.Time `gorm:"column:bid_time;not null;comment:竞拍时间" json:"bid_time"`
}
