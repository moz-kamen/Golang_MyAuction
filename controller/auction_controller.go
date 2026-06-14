package controller

import (
	"MyNFT/config"
	"MyNFT/dao"
	"MyNFT/rpc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuctionController struct {
	Config *config.Config
	Rpc    *rpc.BlockRpc
}

func NewAuctionController(config *config.Config) *BlockController {
	return &BlockController{
		Config: config,
	}
}

func (controller *BlockController) SelectById(context *gin.Context) {
	// 读取参数
	auctionIdStr := context.Param("auctionId")
	auctionId, err := strconv.ParseInt(auctionIdStr, 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"error": "auctionId格式不正确"})
		return
	}

	// 处理业务
	auction := dao.SelectAuctionById(auctionId)

	// 返回相应
	context.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    auction,
	})
}

func (controller *BlockController) SelectLogListById(context *gin.Context) {
	// 读取参数
	auctionIdStr := context.Param("auctionId")
	auctionId, err := strconv.ParseInt(auctionIdStr, 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"error": "auctionId格式不正确"})
		return
	}

	// 处理业务
	auctionLogList := dao.SelectListByAuctionId(auctionId)

	// 返回相应
	context.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    auctionLogList,
	})
}
