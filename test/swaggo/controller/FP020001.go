package controller

import "github.com/gin-gonic/gin"

// FP020001 godoc
// @Summary     addressing register
// @Description addressing register
// @Tags        paymentSwitch
// @Param       FP020001Request body     models.FP020001I true "addressing register"
// @Success     200                        {object} controller.Response{response=models.FP020001O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/addressing/register [post]
func (c *Controller) FP020001(ctx *gin.Context) {

}
