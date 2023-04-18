package controller

import "github.com/gin-gonic/gin"

// FP010016 godoc
// @Summary     8583 lookup and account verify
// @Description 8583 lookup and account verify
// @Tags        paymentSwitch
// @Param       FP010016Request body     models.FP010016I true "8583 lookup and account verify"
// @Success     200                        {object} controller.Response{response=models.FP010016O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/8583/lookup-and-account-verify [post]
func (c *Controller) FP010016(ctx *gin.Context) {

}
