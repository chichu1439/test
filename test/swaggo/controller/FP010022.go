package controller

import "github.com/gin-gonic/gin"

// FP010022 godoc
// @Summary     20022 sign on
// @Description 20022 sign on
// @Tags        paymentSwitch
// @Param       FP010022Request body     models.FPSSignOn001 true "20022 sign on"
// @Success     200                        {object} controller.Response{response=models.FPSSignOn002}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/20022/sign/on [post]
func (c *Controller) FP010022(ctx *gin.Context) {

}
