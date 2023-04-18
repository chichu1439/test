package controller

import "github.com/gin-gonic/gin"

// FP010001 godoc
// @Summary     20022 credit transfer
// @Description 20022 credit transfer
// @Tags        paymentSwitch
// @Param       FP010001Request body     models.FP010001I true "20022 credit transfer"
// @Success     200                        {object} controller.Response{response=models.FP010001O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/credit-transfer [post]
func (c *Controller) FP010001(ctx *gin.Context) {

}
