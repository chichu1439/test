package controller

import "github.com/gin-gonic/gin"

// FP010023 godoc
// @Summary     8583sign on
// @Description 8583sign on
// @Tags        paymentSwitch
// @Param       FP010023Request body     models.FP0100023I true "8583sign on"
// @Success     200                        {object} controller.Response{response=models.FP0100023I}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/8583/sign/on [post]
func (c *Controller) FP010023(ctx *gin.Context) {

}
