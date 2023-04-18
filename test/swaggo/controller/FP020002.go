package controller

import "github.com/gin-gonic/gin"

// FP020002 godoc
// @Summary     addressing cancel
// @Description addressing cancel
// @Tags        paymentSwitch
// @Param       FP020002Request body     models.FP020002I true "addressing cancel"
// @Success     200                        {object} controller.Response{response=models.FP020002O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/addressing/cancel [post]
func (c *Controller) FP020002(ctx *gin.Context) {

}
