package controller

import "github.com/gin-gonic/gin"

// FP000004 godoc
// @Summary     participant suspension
// @Description participant suspension
// @Tags        paymentSwitch
// @Param       FP000004Request body     models.FP000004I true "participant suspension"
// @Success     200                        {object} controller.Response{response=models.FP000004O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/participant/suspension [post]
func (c *Controller) FP000004(ctx *gin.Context) {

}
