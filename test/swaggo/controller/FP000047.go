package controller

import "github.com/gin-gonic/gin"

// FP000047 godoc
// @Summary     participant update
// @Description participant update
// @Tags        paymentSwitch
// @Param       FP000047Request body     models.FP000047I true "participant update"
// @Success     200                        {object} controller.Response{response=models.FP000047O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/alert/participant/update [post]
func (c *Controller) FP000047(ctx *gin.Context) {

}
