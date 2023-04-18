package controller

import "github.com/gin-gonic/gin"

// FP000014 godoc
// @Summary     verification
// @Description verification
// @Tags        paymentSwitch
// @Param       FP000014Request body     models.FP000014I true "verification"
// @Success     200                        {object} controller.Response{response=models.FP000014O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/random-code/verification [post]
func (c *Controller) FP000014(ctx *gin.Context) {

}
