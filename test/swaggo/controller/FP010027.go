package controller

import "github.com/gin-gonic/gin"

// FP010027 godoc
// @Summary     8583bill-payment
// @Description 8583bill-payment
// @Tags        paymentSwitch
// @Param       FP010027Request body     models.FP010027I true "8583bill-payment"
// @Success     200                        {object} controller.Response{response=models.FP010027O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/8583/bill-payment [post]
func (c *Controller) FP010027(ctx *gin.Context) {

}
