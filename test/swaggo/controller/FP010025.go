package controller

import "github.com/gin-gonic/gin"

// FP010025 godoc
// @Summary     8583sign off
// @Description 8583sign off
// @Tags        paymentSwitch
// @Param       FP010025Request body     models.FP0100023I true "8583sign off"
// @Success     200                        {object} controller.Response{response=models.FP0100023I}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/8583/sign/off [post]
func (c *Controller) FP010025(ctx *gin.Context) {

}
