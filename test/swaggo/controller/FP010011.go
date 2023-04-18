package controller

import "github.com/gin-gonic/gin"

// FP010011 godoc
// @Summary     payment-return detail
// @Description payment-return detail
// @Tags        paymentSwitch
// @Param       FP010011Request body     models.FP010011I true "payment-return detail"
// @Success     200                        {object} controller.Response{response=models.FP010011O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/payment-return/detail [post]
func (c *Controller) FP010011(ctx *gin.Context) {

}
