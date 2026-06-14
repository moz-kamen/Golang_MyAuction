package controller

import (
	"MyNFT/config"
	"MyNFT/rpc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	Config *config.Config
	Rpc    *rpc.TransactionRpc
}

func NewTransactionController(config *config.Config) *TransactionController {
	return &TransactionController{
		Config: config,
		Rpc:    rpc.NewTransactionRpc(config),
	}
}

func (controller *TransactionController) Transactions(context *gin.Context) {
	// 读取参数
	blockHash := context.Param("blockHash")

	// 处理业务
	transactionVoList, err := controller.Rpc.Transactions(blockHash)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回相应
	context.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    transactionVoList,
	})
}

func (controller *TransactionController) TransactionByHash(context *gin.Context) {
	// 读取参数
	transactionHash := context.Param("transactionHash")

	// 处理业务
	transaction, err := controller.Rpc.TransactionByHash(transactionHash)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回相应
	context.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    transaction,
	})
}
