package controller

import "github.com/gin-gonic/gin"

// FP000098 godoc
// @Summary     query redis message
// @Description query redis message
// @Tags        paymentSwitch
// @Param       FP000098Request body     models.FP000098I true "query redis message"
// @Success     200                        {object} controller.Response{response=models.FP000098O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/scenario/query/redis/message [post]
func (c *Controller) FP000098(ctx *gin.Context) {

}
