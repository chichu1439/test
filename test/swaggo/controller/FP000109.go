package controller

import "github.com/gin-gonic/gin"

// FP000109 godoc
// @Summary     adjustment confirm
// @Description adjustment confirm
// @Tags        paymentSwitch
// @Param       FP000109Request body     models.FP000109I true "adjustment confirm"
// @Success     200                        {object} controller.Response{response=models.FP000109O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/balance/adjustment/confirm [post]
func (c *Controller) FP000109(ctx *gin.Context) {

}
