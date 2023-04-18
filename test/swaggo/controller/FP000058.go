package controller

import "github.com/gin-gonic/gin"

// FP000058 godoc
// @Summary     step update
// @Description step update
// @Tags        paymentSwitch
// @Param       FP000058Request body     models.FP000058I true "step update"
// @Success     200                        {object} controller.Response{response=models.FP000058O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/scenario/step/update [post]
func (c *Controller) FP000058(ctx *gin.Context) {

}
