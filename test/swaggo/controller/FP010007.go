package controller

import "github.com/gin-gonic/gin"

// FP010007 godoc
// @Summary     message listin quiry
// @Description message listin quiry
// @Tags        paymentSwitch
// @Param       FP010007Request body     models.FP010007I true "message listin quiry"
// @Success     200                        {object} controller.Response{response=models.FP010007O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/message-list/inquiry [post]
func (c *Controller) FP010007(ctx *gin.Context) {

}
