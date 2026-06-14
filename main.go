package main

import (
	"MyNFT/config"
	"MyNFT/controller"
	"MyNFT/subscriber"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 读取配置
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// 部属接口
	v := r.Group("/api/v1")
	{
		block := v.Group("/blocks")
		{
			// 区块接口
			blockController := controller.NewBlockController(config)

			block.GET("/hash", blockController.BlockHash)
			block.GET("/:blockHash", blockController.BlockByHash)

			// 交易接口
			transactionController := controller.NewTransactionController(config)

			block.GET("/:blockHash/transactions", transactionController.Transactions)
			block.GET("/transactions/:transactionHash", transactionController.TransactionByHash)
		}
	}

	srv := &http.Server{
		Addr:    config.Server.Port,
		Handler: r,
	}
	go func() {
		log.Printf("[MAIN]Web接口服务启动成功,监听端口%s", config.Server.Port)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("[MAIN]Web接口服务启动失败: %v", err)
		}
	}()

	globalCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sign := <-sigChan
		log.Printf("[MAIN]接收到系统信号[%v],正在触发优雅停机...", sign)
		cancel()
	}()

	subscriber.SubscribeAuctionEvents(config, globalCtx)
}
