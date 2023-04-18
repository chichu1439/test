package controller

import "github.com/gin-gonic/gin"

// FP010004 godoc
// @Summary     payment return
// @Description payment return
// @Tags        paymentSwitch
// @Param       FP010004Request body     models.FP010004I true "payment return"
// @Success     200                        {object} controller.Response{response=models.FP010004O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/payment-return [post]
func (c *Controller) FP010004(ctx *gin.Context) {

}
