package entity

import "time"

type Auction struct {
	Id          int64     `gorm:"column:id;primaryKey;comment:主键" json:"id"`
	Seller      string    `gorm:"column:seller;not null;comment:出售者" json:"seller"`
	NftContract string    `gorm:"column:nft_contract;not null;comment:NFT合约地址" json:"nft_contract"`
	NftTokenId  int64     `gorm:"column:nft_token_id;not null;comment:NFT主键" json:"nft_token_id"`
	StartPrice  float64   `gorm:"column:start_price;not null;comment:初始价格" json:"start_price"` // 建议生产环境换成 decimal.Decimal
	StartTime   time.Time `gorm:"column:start_time;not null;comment:拍卖开始时间" json:"start_time"`
	EndTime     time.Time `gorm:"column:end_time;not null;comment:拍卖结束时间" json:"end_time"`
	Status      int8      `gorm:"column:status;not null;default:1;comment:拍卖状态 - 1:拍卖中;2:拍卖结束;3:流拍" json:"status"`
	Winner      *string   `gorm:"column:winner;comment:中拍者" json:"winner"`        // 允许为 NULL，使用指针
	WinPrice    *float64  `gorm:"column:win_price;comment:中拍价格" json:"win_price"` // 允许为 NULL，使用指针
}
