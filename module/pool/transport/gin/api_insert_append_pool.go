package poolgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khoindq/tcbHomework/common"
	poolbiz "github.com/khoindq/tcbHomework/module/pool/biz"
	poolmodel "github.com/khoindq/tcbHomework/module/pool/model"
	poolstorage "github.com/khoindq/tcbHomework/module/pool/storage"
)

// @BasePath /api/v1

// InsertorAppendPool godoc
// @Summary Insert or append a pool
// @Schemes
// @Description Inserts or appends a pool to the database
// @Tags Pool
// @Accept json
// @Produce json
// @Param data body poolmodel.Pool true "Pool object"
// @Success 200 {object} common.successRes
// @Router /pool/insertorappend [post]
func (c *PoolController) InsertOrAppendPoolHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req poolmodel.Pool
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

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
	}
}
