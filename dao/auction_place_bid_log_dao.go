package dao

import (
	"MyNFT/dao/entity"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func SaveAuctionPlaceBidLog(auctionLog entity.AuctionPlaceBidLog) {
	dsn := "xuliang:Aa1243%^*&@tcp(rm-bp13v25t74p555u4igo.mysql.rds.aliyuncs.com:3306)/mynft?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Debug().Create(&auctionLog).Error; err != nil {
		log.Fatal(err)
	}
}
