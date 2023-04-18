package controller

import "github.com/gin-gonic/gin"

// FP010020 godoc
// @Summary     8583echo check
// @Description 8583echo check
// @Tags        paymentSwitch
// @Param       FP010020Request body     models.FP0100023I true "8583echo check"
// @Success     200                        {object} controller.Response{response=models.FP0100023I}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/8583/echo/check [post]
func (c *Controller) FP010020(ctx *gin.Context) {

}
