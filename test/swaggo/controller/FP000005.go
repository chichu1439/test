package controller

import "github.com/gin-gonic/gin"

// FP000005 godoc
// @Summary     participant unsuspension
// @Description participant unsuspension
// @Tags        paymentSwitch
// @Param       FP000005Request body     models.FP000005I true "participant unsuspension"
// @Success     200                        {object} controller.Response{response=models.FP000005O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/participant/unsuspension [post]
func (c *Controller) FP000005(ctx *gin.Context) {

}
