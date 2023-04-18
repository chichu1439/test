package controller

import "github.com/gin-gonic/gin"

// FP010017 godoc
// @Summary     8583 credit transfer
// @Description 8583 credit transfer
// @Tags        paymentSwitch
// @Param       FP010017Request body     models.FP010017I true "8583 credit transfer"
// @Success     200                        {object} controller.Response{response=models.FP010017O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/8583/credit-transfer [post]
func (c *Controller) FP010017(ctx *gin.Context) {

}
