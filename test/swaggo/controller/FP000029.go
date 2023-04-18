package controller

import "github.com/gin-gonic/gin"

// FP000029 godoc
// @Summary     sweeping update
// @Description sweeping update
// @Tags        paymentSwitch
// @Param       FP000029Request body     models.FP000029I true "sweeping update"
// @Success     200                        {object} controller.Response{response=models.FP000029O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/balance/sweeping/update [post]
func (c *Controller) FP000029(ctx *gin.Context) {

}
