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

// GetQuantile godoc
// @Summary GetQuantile of a pool
// @Schemes
// @Description GetQuantile of a pool
// @Tags Pool
// @Accept json
// @Produce json
// @Param data body poolmodel.PoolQuantileGet true "PoolQuantileGet object"
// @Success 200 {object} common.successRes
// @Router /pool/quantile/get [post]
func (c *PoolController) GetQuantileHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req poolmodel.PoolQuantileGet
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		poolStore := poolstorage.NewStore()

		biz := poolbiz.NewGetQuantileBiz(poolStore)

		resp, err := biz.GetQuantile(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(resp, nil, nil))
	}
}
