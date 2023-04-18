package controller

import "github.com/gin-gonic/gin"

// FP020003 godoc
// @Summary     addressing amendment
// @Description addressing amendment
// @Tags        paymentSwitch
// @Param       FP020003Request body     models.FP020003I true "addressing amendment"
// @Success     200                        {object} controller.Response{response=models.FP020003O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/addressing/amendment [post]
func (c *Controller) FP020003(ctx *gin.Context) {

}
