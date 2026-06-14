package dao

import (
	"MyNFT/dao/entity"
	"log"
)
import "gorm.io/driver/mysql"
import "gorm.io/gorm"
import "gorm.io/gorm/schema"

func CreateAuction(auction entity.Auction) {
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

	if err := db.Debug().Create(&auction).Error; err != nil {
		log.Fatal(err)
	}
}

func EndAuction(id int64, winner string, winPrice float64) {
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

	if winner != "0" {
		// 正常结束
		if err := db.Debug().Where("id = ?", id).Updates(entity.Auction{Winner: &winner, WinPrice: &winPrice, Status: 2}).Error; err != nil {
			log.Fatal(err)
		}
	} else {
		// 流拍
		if err := db.Debug().Where("id = ?", id).Updates(entity.Auction{Status: 3}).Error; err != nil {
			log.Fatal(err)
		}
	}
}

func SelectAuctionById(id int64) *entity.Auction {
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

	var auction entity.Auction
	if err := db.Debug().Where("id = ?", id).First(&auction).Error; err != nil {
		log.Fatal(err)
	}

	return &auction
}
