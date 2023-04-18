package controller

import "github.com/gin-gonic/gin"

// FP010021 godoc
// @Summary     echo check
// @Description echo check
// @Tags        paymentSwitch
// @Param       FP010021Request body     models.FPSADMNo005 true "echo check"
// @Success     200                        {object} controller.Response{response=models.FPSEcho006}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/20022/echo/check [post]
func (c *Controller) FP010021(ctx *gin.Context) {

}
