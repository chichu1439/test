package controller

import "github.com/gin-gonic/gin"

// FP000019 godoc
// @Summary     debit credit
// @Description debit credit
// @Tags        paymentSwitch
// @Param       FP000019Request body     models.FP000019I true "debit credit"
// @Success     200                        {object} controller.Response{response=models.FP000019O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/participant-balance/debit-credit [post]
func (c *Controller) FP000019(ctx *gin.Context) {

}
