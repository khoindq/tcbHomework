package poolgin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khoindq/tcbHomework/common"
	poolbiz "github.com/khoindq/tcbHomework/module/pool/biz"
	poolmodel "github.com/khoindq/tcbHomework/module/pool/model"
	poolstorage "github.com/khoindq/tcbHomework/module/pool/storage"
)

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /pool/insertappend [post]

func InsertAppendPoolHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req poolmodel.Pool
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		log.Println(req)
		poolStore := poolstorage.NewStore()

		biz := poolbiz.NewInsertAppendPoolBiz(poolStore)

		status, err := biz.InsertAppendPool(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(map[string]interface{}{
			"status": status,
		}, nil, nil))
		log.Println(common.FakeDB.Pools)
	}
}
