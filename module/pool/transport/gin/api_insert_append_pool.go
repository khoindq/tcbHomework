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
	// InsertOrAppendPoolHandler is a method of the PoolController type.
	// It handles the HTTP request for inserting or appending a pool.

	return func(c *gin.Context) {
		// Anonymous function that serves as the actual handler.

		var req poolmodel.Pool

		// Bind the request body to the req variable.
		// If there's an error in binding, respond with a bad request error.
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		// Create a new instance of the poolstore.
		poolStore := poolstorage.NewStore()

		// Create a new instance of the insertAppendPoolBiz with the poolstore.
		biz := poolbiz.NewInsertAppendPoolBiz(poolStore)

		// Call the InsertAppendPool method of the insertAppendPoolBiz to insert or append the pool.
		status, err := biz.InsertAppendPool(c.Request.Context(), &req)

		// If there's an error in inserting or appending the pool, respond with the error.
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		// Respond with a success response containing the status of the operation.
		c.JSON(http.StatusOK, common.NewSuccessResponse(map[string]interface{}{
			"status": status,
		}, nil, nil))
	}
}
