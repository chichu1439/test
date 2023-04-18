package controller

import "github.com/gin-gonic/gin"

// FP010008 godoc
// @Summary     message-detail inquiry
// @Description message-detail inquiry
// @Tags        paymentSwitch
// @Param       FP010008Request body     models.FP010008I true "message-detail inquiry"
// @Success     200                        {object} controller.Response{response=models.FP010008O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/message-detail/inquiry [post]
func (c *Controller) FP010008(ctx *gin.Context) {

}
