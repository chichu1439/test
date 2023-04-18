package controller

import "github.com/gin-gonic/gin"

// FP000105 godoc
// @Summary     balance adjustment apply
// @Description balance adjustment apply
// @Tags        paymentSwitch
// @Param       FP000105Request body     models.FP000105I true "balance adjustment apply"
// @Success     200                        {object} controller.Response{response=models.FP000105O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/balance/adjustment/apply [post]
func (c *Controller) FP000105(ctx *gin.Context) {

}
