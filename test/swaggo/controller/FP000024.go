package controller

import "github.com/gin-gonic/gin"

// FP000024 godoc
// @Summary     ransaction limit update
// @Description ransaction limit update
// @Tags        paymentSwitch
// @Param       FP000024Request body     models.FP000024I true "ransaction limit update"
// @Success     200                        {object} controller.Response{response=models.FP000024O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/transaction/limit/update [post]
func (c *Controller) FP000024(ctx *gin.Context) {

}
