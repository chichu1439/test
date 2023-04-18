package controller

import "github.com/gin-gonic/gin"

// FP010019 godoc
// @Summary     8583request to pay
// @Description 8583request to pay
// @Tags        paymentSwitch
// @Param       FP010019Request body     models.FP010019I true "8583request to pay"
// @Success     200                        {object} controller.Response{response=models.FP010019O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/8583/request-to-pay [post]
func (c *Controller) FP010019(ctx *gin.Context) {

}
