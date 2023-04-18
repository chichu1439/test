package controller

import "github.com/gin-gonic/gin"

// FP000016 godoc
// @Summary     settlement account balancead justment
// @Description settlement account balancead justment
// @Tags        paymentSwitch
// @Param       FP000016Request body     models.FP000016I true "settlement account balancead justment"
// @Success     200                        {object} controller.Response{response=models.FP000016O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/settlement-account-balance/adjustment [post]
func (c *Controller) FP000016(ctx *gin.Context) {

}
