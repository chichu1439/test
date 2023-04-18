package controller

import "github.com/gin-gonic/gin"

// FP010005 godoc
// @Summary     payment summary
// @Description payment summary
// @Tags        paymentSwitch
// @Param       FP010005Request body     models.FP010005I true "payment summary"
// @Success     200                        {object} controller.Response{response=models.FP010005O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/payment-summary [post]
func (c *Controller) FP010005(ctx *gin.Context) {

}
