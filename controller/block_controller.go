package controller

import (
	"MyNFT/config"
	"MyNFT/rpc"

	"net/http"

	"github.com/gin-gonic/gin"
)

type BlockController struct {
	Config *config.Config
	Rpc    *rpc.BlockRpc
}

func NewBlockController(config *config.Config) *BlockController {
	return &BlockController{
		Config: config,
		Rpc:    rpc.NewBlockRpc(config),
	}
}

func (controller *BlockController) BlockHash(context *gin.Context) {
	// 处理业务
	blockHash, err := controller.Rpc.BlockHash()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回相应
	context.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    blockHash,
	})
}

func (controller *BlockController) BlockByHash(context *gin.Context) {
	// 读取参数
	blockHash := context.Param("blockHash")

	// 处理业务
	block, err := controller.Rpc.BlockByHash(blockHash)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回相应
	context.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    block,
	})
}
