package controller

import "github.com/gin-gonic/gin"

// FP000002 godoc
// @Summary     participant detail inquiry
// @Description participant detail inquiry
// @Tags        paymentSwitch
// @Param       FP000002Request body     models.FP000002I true "participant detail inquiry"
// @Success     200                        {object} controller.Response{response=models.FP000002O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/participant-detail/inquiry [post]
func (c *Controller) FP000002(ctx *gin.Context) {

}
