package controller

import "github.com/gin-gonic/gin"

// FP010014 godoc
// @Summary     payment detail enquiry
// @Description payment detail enquiry
// @Tags        paymentSwitch
// @Param       FP010014Request body     models.FP010014I true "payment detail enquiry"
// @Success     200                        {object} controller.Response{response=models.FP010014O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/payment/detail/enquiry [post]
func (c *Controller) FP010014(ctx *gin.Context) {

}
