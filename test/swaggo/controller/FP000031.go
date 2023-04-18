package controller

import "github.com/gin-gonic/gin"

// FP000031 godoc
// @Summary     monthly report
// @Description monthly report
// @Tags        paymentSwitch
// @Param       FP000031Request body     models.FP000031I true "monthly report"
// @Success     200                        {object} controller.Response{response=models.FP000031O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/banch/daily/and/monthly/report [post]
func (c *Controller) FP000031(ctx *gin.Context) {

}
