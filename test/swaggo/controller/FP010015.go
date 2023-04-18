package controller

import "github.com/gin-gonic/gin"

// FP010015 godoc
// @Summary     20022 account verify v1
// @Description 20022 account verify v1
// @Tags        paymentSwitch
// @Param       FP010015Request body     models.FPSPrxy003 true "20022 account verify v1"
// @Success     200                        {object} controller.Response{response=models.FPSPrxy004}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/iso20022/lookup [post]
func (c *Controller) FP010015(ctx *gin.Context) {

}
