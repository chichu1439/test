package controller

import "github.com/gin-gonic/gin"

// FP000003 godoc
// @Summary     participant amendment
// @Description participant amendment
// @Tags        paymentSwitch
// @Param       FP000003Request body     models.FP000003I true "participant amendment"
// @Success     200                        {object} controller.Response{response=models.FP000003O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/participant/amendment [post]
func (c *Controller) FP000003(ctx *gin.Context) {

}
