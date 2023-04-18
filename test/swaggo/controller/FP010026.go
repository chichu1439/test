package controller

import "github.com/gin-gonic/gin"

// FP010026 godoc
// @Summary     8583credit notification
// @Description 8583credit notification
// @Tags        paymentSwitch
// @Param       FP010026Request body     models.FP0100023I true "8583credit notification"
// @Success     200                        {object} controller.Response{response=models.FP010016O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/8583/credit/notification [post]
func (c *Controller) FP010026(ctx *gin.Context) {

}
