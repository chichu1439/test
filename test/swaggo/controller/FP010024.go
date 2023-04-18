package controller

import "github.com/gin-gonic/gin"

// FP010024 godoc
// @Summary     20022sign off
// @Description 20022sign off
// @Tags        paymentSwitch
// @Param       FP010024Request body     models.FPSSignOff003 true "20022sign off"
// @Success     200                        {object} controller.Response{response=models.FPSSignOff004}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/20022/sign/off [post]
func (c *Controller) FP010024(ctx *gin.Context) {

}
