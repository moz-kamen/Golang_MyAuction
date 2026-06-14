package subscriber

import (
	"MyNFT/config"
	"MyNFT/dao"
	"MyNFT/subscriber/event"
	"MyNFT/util"
	"context"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	TopicCreateAuction = crypto.Keccak256Hash([]byte("CreateAuction(address,uint256,address,uint256,uint256,uint256)"))
	TopicPlaceBid      = crypto.Keccak256Hash([]byte("PlaceBid(address,uint256,uint256)"))
	TopicEndAuction    = crypto.Keccak256Hash([]byte("EndAuction(address,uint256,uint256)"))
	minifiedJSON       = `[
		{"anonymous":false,"inputs":[{"indexed":true,"name":"seller","type":"address"},{"indexed":true,"name":"auctionId","type":"uint256"},{"indexed":false,"name":"nftContract","type":"address"},{"indexed":false,"name":"tokenId","type":"uint256"},{"indexed":false,"name":"startPrice","type":"uint256"},{"indexed":false,"name":"durationHours","type":"uint256"}],"name":"CreateAuction","type":"event"},
		{"anonymous":false,"inputs":[{"indexed":true,"name":"bidder","type":"address"},{"indexed":true,"name":"auctionId","type":"uint256"},{"indexed":false,"name":"amount","type":"uint256"}],"name":"PlaceBid","type":"event"},
		{"anonymous":false,"inputs":[{"indexed":true,"name":"winner","type":"address"},{"indexed":true,"name":"auctionId","type":"uint256"},{"indexed":false,"name":"amount","type":"uint256"}],"name":"EndAuction","type":"event"}
	]`
)

func SubscribeAuctionEvents(config *config.Config, ctx context.Context) {
	auctionAddress := common.HexToAddress(config.EthClient.AuctionAddress)

	// 配置过滤规则
	query := ethereum.FilterQuery{
		Addresses: []common.Address{auctionAddress},
		Topics: [][]common.Hash{
			{TopicCreateAuction, TopicPlaceBid, TopicEndAuction},
		},
	}

	parsedABI, err := abi.JSON(strings.NewReader(minifiedJSON))
	if err != nil {
		log.Fatalf("[ETH]解析合约摘要失败: %v", err)
	}

	var lastProcessedBlock uint64 = 0

	// 消费日志流
	for {
		client, err := ethclient.DialContext(ctx, config.EthClient.WebsocketUrl)
		if err != nil {
			log.Printf("[ETH]获取WebSocket连接失败: %v,5秒后重试...", err)
			time.Sleep(5 * time.Second)
			continue
		}
		log.Printf("[ETH]WebSocket连接成功")

		// 断线数据补偿
		if lastProcessedBlock > 0 {
			log.Printf("[ETH]开始补录断线期间的历史数据,起始区块高度: %d", lastProcessedBlock+1)
			catchUpQuery := query
			catchUpQuery.FromBlock = new(big.Int).SetUint64(lastProcessedBlock + 1)

			// 使用 FilterLogs 同步拉取缺失的日志
			missedLogs, err := client.FilterLogs(ctx, catchUpQuery)
			if err == nil {
				for _, vLog := range missedLogs {
					processAuctionLog(vLog, parsedABI)
					lastProcessedBlock = vLog.BlockNumber
				}
				log.Println("[ETH]断线期间的历史数据已成功补录完毕")
			} else {
				log.Printf("[ETH]断线期间的历史数据补录失败: %v,将直接进入实时监听...", err)
			}
		}

		// 订阅实时日志流
		logsCh := make(chan types.Log)
		sub, err := client.SubscribeFilterLogs(ctx, query, logsCh)
		if err != nil {
			log.Printf("[ETH]实时日志流订阅失败: %v,5秒后重试...", err)
			client.Close()
			time.Sleep(5 * time.Second)
			continue
		}

		log.Println("[ETH]实时日志流订阅成功,开始监听拍卖事件")

		// 消费当前连接的日志流
		isDisconnected := false
		for !isDisconnected {
			select {
			case <-ctx.Done():
				log.Println("[ETH]收到外部退出信号,停止事件监听")
				sub.Unsubscribe()
				client.Close()
				return
			case err := <-sub.Err():
				log.Printf("[ETH]事件监听异常中断: %v", err)
				sub.Unsubscribe()
				client.Close()
				isDisconnected = true // 退出内层循环
			case vLog := <-logsCh:
				if len(vLog.Topics) == 0 {
					continue
				}
				processAuctionLog(vLog, parsedABI)
				lastProcessedBlock = vLog.BlockNumber
			}
		}

		// 延迟缓冲
		time.Sleep(2 * time.Second)
	}
}

func processAuctionLog(vLog types.Log, parsedABI abi.ABI) {
	switch vLog.Topics[0] {
	case TopicCreateAuction:
		handleCreateAuction(vLog, parsedABI)
	case TopicPlaceBid:
		handlePlaceBid(vLog, parsedABI)
	case TopicEndAuction:
		handleEndAuction(vLog, parsedABI)
	}
}

func handleCreateAuction(vLog types.Log, parsedABI abi.ABI) {
	var auctionEvent event.CreateAuctionEvent
	if err := parsedABI.UnpackIntoInterface(&auctionEvent, "CreateAuction", vLog.Data); err == nil {
		indexedMap := make(map[string]interface{})
		if err := abi.ParseTopicsIntoMap(indexedMap, util.FilterIndexedFields(parsedABI.Events["CreateAuction"].Inputs), vLog.Topics[1:]); err == nil {
			auctionEvent.Seller = indexedMap["seller"].(common.Address)
			auctionEvent.AuctionId = indexedMap["auctionId"].(*big.Int)
		}
	}

	// 保存拍卖数据入库
	dao.CreateAuction(auctionEvent.ParseEntity())

	log.Printf("[ETH]handleCreateAuction: %+v", auctionEvent)
}

func handlePlaceBid(vLog types.Log, parsedABI abi.ABI) {
	var auctionEvent event.PlaceBidEvent
	if err := parsedABI.UnpackIntoInterface(&auctionEvent, "PlaceBid", vLog.Data); err == nil {
		indexedMap := make(map[string]interface{})
		if err := abi.ParseTopicsIntoMap(indexedMap, util.FilterIndexedFields(parsedABI.Events["PlaceBid"].Inputs), vLog.Topics[1:]); err == nil {
			auctionEvent.Bidder = indexedMap["bidder"].(common.Address)
			auctionEvent.AuctionId = indexedMap["auctionId"].(*big.Int)
		}
	}

	// 保存拍卖竞拍数据入库
	dao.SaveAuctionPlaceBidLog(auctionEvent.ParseEntity())

	log.Printf("[ETH]handlePlaceBid: %+v", auctionEvent)
}

func handleEndAuction(vLog types.Log, parsedABI abi.ABI) {
	var auctionEvent event.EndAuctionEvent
	if err := parsedABI.UnpackIntoInterface(&auctionEvent, "EndAuction", vLog.Data); err == nil {
		indexedMap := make(map[string]interface{})
		if err := abi.ParseTopicsIntoMap(indexedMap, util.FilterIndexedFields(parsedABI.Events["EndAuction"].Inputs), vLog.Topics[1:]); err == nil {
			auctionEvent.Winner = indexedMap["winner"].(common.Address)
			auctionEvent.AuctionId = indexedMap["auctionId"].(*big.Int)
		}
	}

	// 结束竞拍
	winPrice, _ := auctionEvent.Amount.Float64()
	dao.EndAuction(auctionEvent.AuctionId.Int64(), auctionEvent.Winner.Hex(), winPrice)

	log.Printf("[ETH]handleEndAuction: %+v", auctionEvent)
}
