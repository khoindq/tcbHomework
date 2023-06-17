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
	// GetQuantileHandler is a method of the PoolController type.
	// It handles the HTTP request for retrieving a quantile value.

	return func(c *gin.Context) {
		// Anonymous function that serves as the actual handler.

		var req poolmodel.PoolQuantileGet

		// Bind the request body to the req variable.
		// If there's an error in binding, respond with a bad request error.
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		// Create a new instance of the poolstore.
		poolStore := poolstorage.NewStore()

		// Create a new instance of the getQuantileBiz with the poolstore.
		biz := poolbiz.NewGetQuantileBiz(poolStore)

		// Call the GetQuantile method of the getQuantileBiz to retrieve the quantile value.
		resp, err := biz.GetQuantile(c.Request.Context(), &req)

		// If there's an error in retrieving the quantile, respond with the error.
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		// Respond with a success response containing the quantile value.
		c.JSON(http.StatusOK, common.NewSuccessResponse(resp, nil, nil))
	}
}
