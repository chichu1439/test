package controller

import "github.com/gin-gonic/gin"

// FP010006 godoc
// @Summary     credit-transfer detail
// @Description credit-transfer detail
// @Tags        paymentSwitch
// @Param       FP010006Request body     models.FP010006I true "credit-transfer detail"
// @Success     200                        {object} controller.Response{response=models.FP010006O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/credit-transfer/detail [post]
func (c *Controller) FP010006(ctx *gin.Context) {

}
