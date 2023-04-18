package controller

import "github.com/gin-gonic/gin"

// FP000112 godoc
// @Summary     payment enum list
// @Description payment enum list
// @Tags        paymentSwitch
// @Param       FP000112Request body     models.FP000112I true "payment enum list"
// @Success     200                        {object} controller.Response{response=models.FP000112O}
// @Failure     500                {object} controller.Error
// @Router     /v3/nfps/payment/enum/list [post]
func (c *Controller) FP000112(ctx *gin.Context) {

}
