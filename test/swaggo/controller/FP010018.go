package controller

import "github.com/gin-gonic/gin"

// FP010018 godoc
// @Summary     request to pay
// @Description request to pay
// @Tags        paymentSwitch
// @Param       FP010018Request body     models.FPSPain013 true "request to pay"
// @Success     200                        {object} controller.Response{response=models.FPSPacs002}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/request/to/pay [post]
func (c *Controller) FP010018(ctx *gin.Context) {

}
