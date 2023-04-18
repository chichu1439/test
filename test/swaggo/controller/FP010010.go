package controller

import "github.com/gin-gonic/gin"

// FP010010 godoc
// @Summary     direct-debit detail
// @Description direct-debit detail
// @Tags        paymentSwitch
// @Param       FP010010Request body     models.FP010010I true "direct-debit detail"
// @Success     200                        {object} controller.Response{response=models.FP010010O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/direct-debit/detail [post]
func (c *Controller) FP010010(ctx *gin.Context) {

}
