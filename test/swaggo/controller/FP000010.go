package controller

import "github.com/gin-gonic/gin"

// FP000010 godoc
// @Summary     other participant detail inquiry
// @Description other participant detail inquiry
// @Tags        paymentSwitch
// @Param       FP000010Request body     models.FP000010I true "other participant detail inquiry"
// @Success     200                        {object} controller.Response{response=models.FP000010O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/other-participant-detail/inquiry [post]
func (c *Controller) FP000010(ctx *gin.Context) {

}
