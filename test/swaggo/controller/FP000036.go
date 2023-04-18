package controller

import "github.com/gin-gonic/gin"

// FP000036 godoc
// @Summary     balance alter noti
// @Description balance alter noti
// @Tags        paymentSwitch
// @Param       FP000036Request body     models.FP000036I true "balance alter noti"
// @Success     200                        {object} controller.Response{response=models.FP000036O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/balance/alter/noti [post]
func (c *Controller) FP000036(ctx *gin.Context) {

}
