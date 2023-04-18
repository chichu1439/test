package controller

import "github.com/gin-gonic/gin"

// FP000040 godoc
// @Summary     prodfee update
// @Description prodfee update
// @Tags        paymentSwitch
// @Param       FP000040Request body     models.FP000040I true "prodfee update"
// @Success     200                        {object} controller.Response{response=models.FP000040O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/prod/fee/update [post]
func (c *Controller) FP000040(ctx *gin.Context) {

}
